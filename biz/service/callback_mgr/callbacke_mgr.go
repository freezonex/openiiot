package callback_mgr

import (
	"context"
	"fmt"
	"time"
)

type GetUserByTokenDBFunc func(ctx context.Context, usertoken string) (*time.Time, string, error)

var callBackMap map[string]GetUserByTokenDBFunc

func init() {
	callBackMap = make(map[string]GetUserByTokenDBFunc)
}

func RegisterCallBack(key string, callBack GetUserByTokenDBFunc) {
	callBackMap[key] = callBack
}

func CallUserByTokenDBFunc(key string, ctx context.Context, usertoken string) (*time.Time, string, error) {
	if callBack, ok := callBackMap[key]; ok {
		return callBack(ctx, usertoken)
	}
	return nil, "", fmt.Errorf("callBack(%s) not found", key)
}
