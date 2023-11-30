package common

import (
	"testing"

	"freezonex/openiiot/biz/service/utils/testtypes"
	"github.com/stretchr/testify/assert"
)

func TestUseValueOrFallback(t *testing.T) {
	tests := []struct {
		testtypes.Label
		value     int
		fallback  int
		predicate func(int) bool
		expected  int
	}{
		{
			Label: testtypes.Label{
				Name: "true evaluates to value",
			},
			value:    1,
			fallback: 2,
			predicate: func(i int) bool {
				return i%2 == 1
			},
			expected: 1,
		},
		{
			Label: testtypes.Label{
				Name: "false evaluates to fallback",
			},
			value:    1,
			fallback: 2,
			predicate: func(i int) bool {
				return i%2 == 0
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			output := UseValueOrFallback(test.value, test.fallback, test.predicate)
			assert.Equal(t, test.expected, output)
		})
	}
}

func TestIdentity(t *testing.T) {
	assert.Equal(t, 1, Identity(1))
	assert.Equal(t, "", Identity(""))
}
