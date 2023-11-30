package supos

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	logger "github.com/sirupsen/logrus"
)

// DoRequest performs an openAPI request with AK/SK authorization.
func (a *SuposService) DoRequest(suposApiEntity SuposApiEntity, jsonBody string, header map[string]string, pathParam ...interface{}) (*ApiResponse, error) {
	t1 := time.Now()
	apiURL := fmt.Sprintf(suposApiEntity.URL, pathParam...)

	wholeURL := a.c.URL + apiURL
	logger.Printf("Beginning openAPI request URL: %s, method: %s, body: %s", wholeURL, suposApiEntity.Method, jsonBody)

	req, err := http.NewRequest(suposApiEntity.Method, wholeURL, strings.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	headers := a.getSignatureHeader(jsonBody, suposApiEntity.Method, apiURL)
	for k, v := range header {
		headers[k] = v
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if suposApiEntity.Method == "GET" || suposApiEntity.Method == "DELETE" {
		if jsonBody != "" {
			// Parse the JSON body into a map
			var queryMap map[string]interface{}
			err = json.Unmarshal([]byte(jsonBody), &queryMap)
			if err != nil {
				return nil, err
			}
			// Add parsed parameters to the query string
			q := req.URL.Query()
			for key, value := range queryMap {
				q.Add(key, fmt.Sprintf("%v", value))
			}
			req.URL.RawQuery = q.Encode()
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	logger.Printf("OpenAPI call ended with HttpCode: %d, result: %s, time taken: %vms", resp.StatusCode, string(body), time.Since(t1).Milliseconds())

	return &ApiResponse{
		HttpCode:     resp.StatusCode,
		ResponseBody: string(body),
	}, nil
}

func (a *SuposService) getSignatureHeader(requestBody, method, apiURL string) map[string]string {
	headers := make(map[string]string)
	signature := a.buildSignature(method, apiURL, requestBody)
	authHeader := fmt.Sprintf("Sign %s-%s", a.c.AK, signature)

	//Sign c01f69318340143ce7cc9c3870173fa0-006aab2fe6e5203a2afafbc07e02a9b43f19d664b99487cc4e80907c2775fc78
	headers["Authorization"] = authHeader
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	logger.Printf("AK/SK Signature: %s", authHeader)

	return headers
}

func (a *SuposService) buildSignature(method, apiURL, requestBody string) string {
	var sb strings.Builder

	sb.WriteString(method + "\n")
	sb.WriteString(apiURL + "\n")
	sb.WriteString("application/json" + "\n")

	if method == "GET" || method == "DELETE" {
		var queryMap map[string]interface{}
		json.Unmarshal([]byte(requestBody), &queryMap)
		finalSignature := beanToQueryString(queryMap, false)
		queryString := getSortQueryStr(finalSignature)
		sb.WriteString(queryString)
	}

	sb.WriteString("\n")
	sb.WriteString("\n")
	//sb="POST\n/open-api/auth/v2/oauth2/token\napplication/json\n\n\n"
	//sb="GET\n/open-api/auth/v2/users/hongzhi\napplication/json\n\n\n"

	mac := hmac.New(sha256.New, []byte(a.c.SK))
	mac.Write([]byte(sb.String()))
	signature := hex.EncodeToString(mac.Sum(nil))

	//006aab2fe6e5203a2afafbc07e02a9b43f19d664b99487cc4e80907c2775fc78
	//05ddc16fb449316f6ac5bd7556160510c7a120632bf4a8c2eebafeffb93a3b94
	return signature
}

// getSortQueryStr sorts the query parameters of a given API path
func getSortQueryStr(apiPath string) string {
	if apiPath == "" {
		return ""
	}

	params := getUrlParams(apiPath)
	keys := make([]string, 0, len(params))
	lowerToOriginalKeyMap := make(map[string]string)
	for k := range params {
		lowerKey := strings.ToLower(k)
		keys = append(keys, lowerKey)
		lowerToOriginalKeyMap[lowerKey] = k
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		originalKey := lowerToOriginalKeyMap[k]
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(params[originalKey])
		sb.WriteString("&")
	}

	sortedQuery := sb.String()
	sortedQuery = strings.TrimSuffix(sortedQuery, "&")

	return sortedQuery
}

// getUrlParams parses the URL and returns a map of query parameters
func getUrlParams(urlStr string) map[string]string {
	result := make(map[string]string)
	parsedUrl, err := url.Parse("http://dummy.com/?" + urlStr)
	if err != nil {
		return result
	}

	params := parsedUrl.Query()
	for k := range params {
		result[k] = params.Get(k)
	}

	return result
}

func beanToQueryString(paramMap map[string]interface{}, isEncoder bool) string {
	if len(paramMap) == 0 {
		return ""
	}

	var sb strings.Builder
	for key, value := range paramMap {
		if isEncoder {
			sb.WriteString(url.QueryEscape(key))
			sb.WriteString("=")
			sb.WriteString(url.QueryEscape(fmt.Sprintf("%v", value)))
		} else {
			sb.WriteString(key)
			sb.WriteString("=")
			sb.WriteString(fmt.Sprintf("%v", value))
		}
		sb.WriteString("&")
	}

	queryString := sb.String()
	return strings.TrimSuffix(queryString, "&")
}
