package error_utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	errorOne   = errors.New("a")
	errorTwo   = errors.New("b")
	errorThree = errors.New("c")
)

func TestAnyErrorMatch(t *testing.T) {
	assert.Equal(t, false, AnyErrorMatch(errorOne))
	assert.Equal(t, false, AnyErrorMatch(errorOne, errorTwo))
	assert.Equal(t, true, AnyErrorMatch(errorOne, errorTwo, errorOne))
	assert.Equal(t, true, AnyErrorMatch(errorOne, errorOne, errorOne))
}
