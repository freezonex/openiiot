package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{1, 3, 2, 0, -1},
			expected: []int{1, 3, 2, 0, -1},
		},
	}

	for _, test := range tests {
		iterator := FromSlice(test.input)
		elem := iterator()
		index := 0
		for elem.IsSome() {
			assert.Equal(t, test.expected[index], elem.Unwrap())
			elem = iterator()
			index++
		}
		assert.Equal(t, len(test.expected), index)
	}
}

func TestCollectToSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{1, 3, 2, 0, -1},
			expected: []int{1, 3, 2, 0, -1},
		},
	}

	for _, test := range tests {
		iterator := FromSlice(test.input)
		collected := CollectToSlice(iterator)
		assert.Equal(t, test.expected, collected)
	}
}
