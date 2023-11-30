package binding

import (
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/bytedance/go-tagexpr/binding"
	"github.com/bytedance/go-tagexpr/binding/gjson"
)

// BindAndValidate binds data from *app.RequestContext to obj and validates them if needed.
func BindAndValidate(ctx *app.RequestContext, obj interface{}) error {
	return binding.Default().IBindAndValidate(obj, WrapRequest(ctx), ctx.Params)
}

// WriteHeader writes header and cookie to response,
// according to the 'header' and 'cookie' struct tags.
func WriteHeader(ctx *app.RequestContext, result interface{}) error {
	resp := wrapRespond(ctx)
	err := defaultResponding.WriteHeader(resp, result)
	if err == nil {
		for key, values := range resp.Header() {
			for _, value := range values {
				ctx.Response.Header.Add(key, value)
			}
		}
	}
	return err
}

// UseStdJSONUnmarshaler uses encoding/json as json libary
// NOTE: 当前版本默认使用encoding/json，无需额外设置
func UseStdJSONUnmarshaler() {
	binding.ResetJSONUnmarshaler(json.Unmarshal)
}

// UseGJSONUnmarshaler uses github.com/bytedance/go-tagexpr/v2/binding/gjson as json libary
// NOTE: 老版本默认使用gjson，支持string与int兼容的同时，也存在与标准包不完全一致问题
func UseGJSONUnmarshaler() {
	gjson.UseJSONUnmarshaler()
}

// UseThirdPartyJSONUnmarshaler uses third-party json libary for binding
func UseThirdPartyJSONUnmarshaler(unmarshaler func(data []byte, v interface{}) error) {
	binding.ResetJSONUnmarshaler(unmarshaler)
}
