package binding

import (
	"net/http"
	"net/url"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/bytedance/go-tagexpr/v2/binding"
)

func WrapRequest(ctx *app.RequestContext) binding.Request {
	r := &bindRequest{
		ctx: ctx,
	}
	return r
}

type bindRequest struct {
	ctx *app.RequestContext
}

func (r *bindRequest) GetQuery() url.Values {
	queryMap := make(url.Values)
	r.ctx.QueryArgs().VisitAll(func(key, value []byte) {
		keyStr := string(key)
		values, _ := queryMap[keyStr]
		values = append(values, string(value))
		queryMap[keyStr] = values
	})
	return queryMap
}

func (r *bindRequest) GetPostForm() (url.Values, error) {
	postMap := make(url.Values)
	r.ctx.PostArgs().VisitAll(func(key, value []byte) {
		keyStr := string(key)
		values, _ := postMap[keyStr]
		values = append(values, string(value))
		postMap[keyStr] = values
	})
	mf, err := r.ctx.MultipartForm()
	if err == nil && mf.Value != nil {
		for k, v := range mf.Value {
			if len(v) > 0 {
				postMap[k] = v
			}
		}
	}
	return postMap, nil
}

func (r *bindRequest) GetForm() (url.Values, error) {
	formMap := make(url.Values)
	r.ctx.QueryArgs().VisitAll(func(key, value []byte) {
		keyStr := string(key)
		values, _ := formMap[keyStr]
		values = append(values, string(value))
		formMap[keyStr] = values
	})
	r.ctx.PostArgs().VisitAll(func(key, value []byte) {
		keyStr := string(key)
		values, _ := formMap[keyStr]
		values = append(values, string(value))
		formMap[keyStr] = values
	})
	return formMap, nil
}

func (r *bindRequest) GetCookies() []*http.Cookie {
	var cookies []*http.Cookie
	r.ctx.Request.Header.VisitAllCookie(func(key, value []byte) {
		cookies = append(cookies, &http.Cookie{
			Name:  string(key),
			Value: string(value),
		})
	})
	return cookies
}

func (r *bindRequest) GetHeader() http.Header {
	header := make(http.Header)
	r.ctx.Request.Header.VisitAll(func(key, value []byte) {
		keyStr := string(key)
		values, _ := header[keyStr]
		values = append(values, string(value))
		header[keyStr] = values
	})
	return header
}

func (r *bindRequest) GetMethod() string {
	return b2s(r.ctx.Request.Method())
}

func (r *bindRequest) GetContentType() string {
	return b2s(r.ctx.Request.Header.ContentType())
}

func (r *bindRequest) GetBody() ([]byte, error) {
	return r.ctx.Request.Body(), nil
}
