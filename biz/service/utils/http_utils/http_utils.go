package http_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	//"moul.io/http2curl"

	"freezonex/openiiot/biz/service/utils/error_utils"
	"freezonex/openiiot/biz/service/utils/monad/result"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type Query struct {
	Key, Value string
}
type Path struct {
	Key, Value string
}

type Header struct {
	Key, Value string
}

// generic http request, response body is returned
func SendHttpRequest(ctx context.Context, method, url string, data interface{}, querys []Query, paths []Path, header []Header) ([]byte, error) {
	for _, path := range paths {
		url = strings.Replace(url, fmt.Sprintf(":%v", path.Key), path.Value, 1)
	}
	p := ""
	if len(querys) != 0 {
		url = url + "?"
		for i, query := range querys {
			p += query.Key + "=" + query.Value
			if i < len(querys)-1 {
				p += "&"
			}
		}
	}
	url += p
	bytesData, _ := json.Marshal(data)
	requestBody := strings.NewReader(string(bytesData))
	logs.CtxInfof(ctx, "NewHttpRequest[%v] data: %v", method, data)
	logs.CtxInfof(ctx, "NewHttpRequest[%v] requestBody: %v", method, requestBody)
	logs.CtxInfof(ctx, "NewHttpRequest[%v] url: %v", method, url)
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		logs.CtxErrorf(ctx, "SendHttpRequest[%v] failed, err: %v", url, err)
		return nil, fmt.Errorf("SendHttpRequest[%v] failed, err: %v", url, err)
	}
	req.Header.Add("Content-Type", "application/json")
	for _, h := range header {
		req.Header.Add(h.Key, h.Value)
	}
	// Determine whether to skip SSL verification based on the URL scheme
	insecureSkipVerify := req.URL.Scheme == "https"
	command, err := GetCurlCommand(req, insecureSkipVerify)
	if err != nil {
		logs.CtxWarnf(ctx, "Couldn't generate curl")
	} else {
		logs.CtxInfof(ctx, "Generated curl: %s", command)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logs.CtxErrorf(ctx, "SendHttpRequest[%v] failed, err: %v", url, err)
		return nil, fmt.Errorf("SendHttpRequest[%v] failed, err: %v", url, err)
	}
	if resp.StatusCode != 200 {
		logs.CtxWarnf(ctx, "status code isnt 200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.CtxErrorf(ctx, "SendHttpRequest[%v] read response body failed, err: %v", url, err)
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			logs.CtxWarnf(ctx, "SendHttpRequest[%v] close response body err: %v", url, closeErr.Error())
		}
	}()
	return body, nil
}

func Get(ctx context.Context, url string, querys []Query, paths []Path, header []Header) ([]byte, error) {
	return SendHttpRequest(ctx, "GET", url, "", querys, paths, header)
}

func Delete(ctx context.Context, url string, querys []Query, paths []Path, header []Header) ([]byte, error) {
	return SendHttpRequest(ctx, "DELETE", url, "", querys, paths, header)
}

func GetWithUnmarshal(ctx context.Context, ret interface{}, url string, querys []Query, paths []Path, header []Header) error {
	if data, err := SendHttpRequest(ctx, "GET", url, "", querys, paths, header); err != nil {
		return fmt.Errorf("GetWithUnmarshal[%v] SendHttpRequest err: %v", url, err)
	} else if err = json.Unmarshal(data, ret); err != nil {
		return fmt.Errorf("GetWithUnmarshal[%v] Unmarshal data: %v err: %v", url, string(data), err)
	} else {
		return nil
	}
}

func Post(ctx context.Context, url string, data interface{}, querys []Query, paths []Path, header []Header) ([]byte, error) {
	return SendHttpRequest(ctx, "POST", url, data, querys, paths, header)
}

func PostWithUnmarshal(ctx context.Context, ret interface{}, url string, datas interface{}, querys []Query, paths []Path, header []Header) error {
	if data, err := SendHttpRequest(ctx, "POST", url, datas, querys, paths, header); err != nil {
		return fmt.Errorf("PostWithUnmarshal[%v] SendHttpRequest err: %v", url, err)
	} else if err = json.Unmarshal(data, ret); err != nil {
		return fmt.Errorf("PostWithUnmarshal[%v] Unmarshal data: %v err: %v", url, string(data), err)
	} else {
		return nil
	}
}

// ParseForm returns the bytes of the first file matching the form key name in a request context.
func ParseForm(reqCtx *app.RequestContext, name string) result.Result[[]byte] {
	fileHeader, err := reqCtx.FormFile(name)
	if err != nil {
		return result.Err[[]byte](err)
	}
	return ParseMultipartFileHeader(fileHeader)
}

// ParseMultipartFileHeader parses a multipart.FileHeader into the bytes of the file associated with the FileHeader.
func ParseMultipartFileHeader(fileHeader *multipart.FileHeader) result.Result[[]byte] {
	if fileHeader == nil {
		return result.Err[[]byte](error_utils.ErrNull)
	}
	file, err := fileHeader.Open()
	if err != nil {
		return result.Err[[]byte](err)
	}
	return ParseMultipartFile(file)
}

// ParseMultipartFile parses a multipart.File into the bytes that represent the file.
func ParseMultipartFile(file multipart.File) result.Result[[]byte] {
	data, err := io.ReadAll(file)
	if err != nil {
		return result.Err[[]byte](err)
	}
	return result.Ok(data)
}
