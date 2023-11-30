package error_utils

import (
	"errors"

	"freezonex/openiiot/biz/service/utils/iter"
)

var (
	ErrNull = errors.New("null value")
)

// AnyErrorMatch checks if error matches on of the supplied errors.
func AnyErrorMatch(err error, matches ...error) bool {
	return iter.Any(iter.From(matches...), func(e error) bool {
		return errors.Is(err, e)
	})
}
