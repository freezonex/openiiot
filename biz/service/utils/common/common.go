package common

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"freezonex/openiiot/biz/model/base_req"

	"gorm.io/gen/field"

	"github.com/google/uuid"

	"freezonex/openiiot/biz/config/consts"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

func ProtectNullStringPointer(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ProtectNullBoolPointer(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func GetTimeStringFromTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.String()
}

func GetUUID() int64 {
	id := uuid.New()
	a := binary.BigEndian.Uint64(id[:16])
	if a > 9223372036854775807 {
		return GetUUID()
	}
	return int64(a)
}

// UseValueOrFallback returns value if predicate passes, if not, fallback value is returned.
func UseValueOrFallback[A any](value, fallback A, predicate func(A) bool) A {
	if predicate(value) {
		return value
	} else {
		return fallback
	}
}

func ConvertStringSliceToInt64Slice(stringSlice []string) []int64 {
	var intSlice []int64

	for _, value := range stringSlice {
		intValue, _ := strconv.ParseInt(value, 10, 64)
		intSlice = append(intSlice, intValue)
	}

	return intSlice
}

func SliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Identity function returns its argument.
func Identity[A any](a A) A {
	return a
}

type fieldGetter interface {
	GetFieldByName(string) (field.OrderExpr, bool)
}

func GetSortingField(ctx context.Context, b fieldGetter, Pagination *base_req.PaginationRequest) (*field.Expr, error) {
	orderExp, _ := b.GetFieldByName("created_at")
	if Pagination != nil && Pagination.SortKey != "" {
		ok := false
		orderExp, ok = b.GetFieldByName(Pagination.SortKey)
		if !ok {
			logs.CtxErrorf(ctx, "requested sorting key not found, key: %s", Pagination.SortKey)
			return nil, consts.ErrKeyNotFound
		}
	}
	var orderBy field.Expr
	if Pagination != nil && Pagination.SortDescending {
		orderBy = orderExp.Desc()
	} else {
		orderBy = orderExp
	}
	return &orderBy, nil
}

func PrettyPrint(i interface{}) {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		fmt.Printf("%+v\n", i)
	} else {
		fmt.Println(string(s))
	}
}

func RemoveEmailSuffix(email string) string {
	return strings.Replace(strings.ToLower(email), "@freezonex.io", "", -1)
}

func IsOpenChatId(id string) bool {
	return strings.HasPrefix(id, "oc_")
}

func GetRuntimeIdc() string {
	return os.Getenv("RUNTIME_IDC_NAME")
}

// Int64ToString converts an int64 to a string.
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// StringToInt64 converts a string to an int64. It returns 0 if the conversion is not successful.
func StringToInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}
