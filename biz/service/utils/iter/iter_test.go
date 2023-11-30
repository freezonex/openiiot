package iter

import (
	"fmt"
	"strconv"
	"testing"

	"freezonex/openiiot/biz/service/utils/monad/option"
	"freezonex/openiiot/biz/service/utils/pair"
	p "freezonex/openiiot/biz/service/utils/pair"
	"freezonex/openiiot/biz/service/utils/testtypes"

	"github.com/stretchr/testify/assert"
)

func TestFrom(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    []int
		expected func() option.Option[int]
	}{
		{
			Label: testtypes.Label{
				Name: "empty input should return empty iterator",
			},
			input:    []int{},
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "single input returns a singleton iterator",
			},
			input:    []int{1},
			expected: From(1),
		},
		{
			Label: testtypes.Label{
				Name: "multiple inputs returns an iterator of all inputs",
			},
			input:    []int{1, 3, 2, 0, -1},
			expected: From(1, 3, 2, 0, -1),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			output := From(test.input...)
			assert.Equal(t, CollectToSlice(test.expected), CollectToSlice(output))
		})
	}
}

func TestEnumerate(t *testing.T) {
	tests := []struct {
		testtypes.Label
		iter     func() option.Option[int]
		expected func() option.Option[pair.Pair[uint, int]]
	}{
		{
			Label: testtypes.Label{
				Name: "enumerate empty list",
			},
			iter:     From[int](),
			expected: From[pair.Pair[uint, int]](),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(
				t,
				CollectToSlice(test.expected),
				CollectToSlice(Enumerate(test.iter)))
		})
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    func() option.Option[int]
		fn       func(int) string
		expected func() option.Option[string]
	}{
		{
			Label: testtypes.Label{
				Name: "empty input returns empty output",
			},
			input: From[int](),
			fn: func(i int) string {
				return fmt.Sprint(i * 2)
			},
			expected: From[string](),
		},
		{
			Label: testtypes.Label{
				Name: "singleton iterator maps correctly",
			},
			input: From(1),
			fn: func(i int) string {
				return fmt.Sprint(i * 2)
			},
			expected: From("2"),
		},
		{
			Label: testtypes.Label{
				Name: "iterator maps correctly",
			},
			input: From(1, 3, 2, 0, -1),
			fn: func(i int) string {
				return fmt.Sprint(i * 2)
			},
			expected: From("2", "6", "4", "0", "-2"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Map(test.input, test.fn)
			assert.Equal(t, CollectToSlice(test.expected), CollectToSlice(out))
		})
	}

}

func TestFilter(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    func() option.Option[int]
		fn       func(int) bool
		expected func() option.Option[int]
	}{
		{
			Label: testtypes.Label{
				Name: "empty iterator filters to empty iterator",
			},
			input: From[int](),
			fn: func(_ int) bool {
				return true
			},
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "true function yields all elements in iterator",
			},
			input: From(1, 2, 3),
			fn: func(_ int) bool {
				return true
			},
			expected: From(1, 2, 3),
		},
		{
			Label: testtypes.Label{
				Name: "false function yields no elements in iterator",
			},
			input: From(1, 2, 3),
			fn: func(_ int) bool {
				return false
			},
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "filter only yields elements that results in true",
			},
			input: From(1, 2, 3),
			fn: func(i int) bool {
				return i%2 != 0
			},
			expected: From(1, 3),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Filter(test.input, test.fn)
			assert.Equal(t, CollectToSlice(test.expected), CollectToSlice(out))
		})
	}
}

func TestFold(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    func() option.Option[int64]
		initial  string
		fn       func(string, int64) string
		expected string
	}{
		{
			Label: testtypes.Label{
				Name: "empty iterator yields initial value",
			},
			input:   From[int64](),
			initial: "a",
			fn: func(s string, i int64) string {
				return s + " " + strconv.FormatInt(i, 10)
			},
			expected: "a",
		},
		{
			Label: testtypes.Label{
				Name: "iterator yields correct result value from accumulator function",
			},
			input:   From[int64](1, 2, 3),
			initial: "a",
			fn: func(s string, i int64) string {
				return s + " " + strconv.FormatInt(i, 10)
			},
			expected: "a 1 2 3",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Fold(test.input, test.initial, test.fn)
			assert.Equal(t, test.expected, out)
		})
	}
}

func TestUnfold(t *testing.T) {
	tests := []struct {
		testtypes.Label
		seed     int
		fn       func(int) int
		expected func() option.Option[int]
		limit    uint
	}{
		{
			Label: testtypes.Label{
				Name: "unfold should produce correct order of elements",
			},
			seed: 0,
			fn: func(i int) int {
				return i + 1
			},
			expected: From(0, 1, 2, 3, 4),
			limit:    5,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Unfold(test.seed, test.fn)
			assert.Equal(t, CollectToSlice(test.expected), CollectToSlice(Take(out, test.limit)))
		})
	}
}

func TestZip(t *testing.T) {
	tests := []struct {
		testtypes.Label
		firstIter  func() option.Option[int]
		secondIter func() option.Option[string]
		expected   func() option.Option[p.Pair[int, string]]
	}{
		{
			Label: testtypes.Label{
				Name: "empty iterators yield empty iterator",
			},
			firstIter:  From[int](),
			secondIter: From[string](),
			expected:   From[p.Pair[int, string]](),
		},
		{
			Label: testtypes.Label{
				Name: "first empty iterator yield empty iterator",
			},
			firstIter:  From[int](),
			secondIter: From("a", "b", "c"),
			expected:   From[p.Pair[int, string]](),
		},
		{
			Label: testtypes.Label{
				Name: "second empty iterator yield empty iterator",
			},
			firstIter:  From(0, 1, 2),
			secondIter: From[string](),
			expected:   From[p.Pair[int, string]](),
		},
		{
			Label: testtypes.Label{
				Name: "truncate to shortest iterator (first)",
			},
			firstIter:  From(0, 1, 2, 3, 4, 5),
			secondIter: From("a", "b", "c"),
			expected:   From(p.New(0, "a"), p.New(1, "b"), p.New(2, "c")),
		},
		{
			Label: testtypes.Label{
				Name: "truncate to shortest iterator (second)",
			},
			firstIter:  From(0, 1, 2),
			secondIter: From("a", "b", "c", "d", "e"),
			expected:   From(p.New(0, "a"), p.New(1, "b"), p.New(2, "c")),
		},
		{
			Label: testtypes.Label{
				Name: "truncate to shortest iterator (second)",
			},
			firstIter:  From(0, 1, 2, 3, 4),
			secondIter: From("a", "b", "c", "d", "e"),
			expected:   From(p.New(0, "a"), p.New(1, "b"), p.New(2, "c"), p.New(3, "d"), p.New(4, "e")),
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Zip(test.firstIter, test.secondIter)
			assert.Equal(t, CollectToSlice(test.expected), CollectToSlice(out))
		})
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    func() option.Option[int]
		fn       func(int) bool
		expected bool
	}{
		{
			Label: testtypes.Label{
				Name: "empty iterator always returns true",
			},
			input: From[int](),
			fn: func(_ int) bool {
				return false
			},
			expected: true,
		},
		{
			Label: testtypes.Label{
				Name: "if there is any false evaluation, return false",
			},
			input: From(1, 2, 3),
			fn: func(i int) bool {
				return i%2 == 0
			},
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "if all evaluations are true, return true",
			},
			input: From(1, 2, 3),
			fn: func(i int) bool {
				return i < 10
			},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := All(test.input, test.fn)
			assert.Equal(t, test.expected, out)
		})
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    func() option.Option[int]
		fn       func(int) bool
		expected bool
	}{
		{
			Label: testtypes.Label{
				Name: "empty iterator always returns false",
			},
			input: From[int](),
			fn: func(_ int) bool {
				return true
			},
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "if there is no true evaluation, return false",
			},
			input: From(1, 2, 3),
			fn: func(i int) bool {
				return i > 5
			},
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "if any evaluations are true, return true",
			},
			input: From(1, 2, 3),
			fn: func(i int) bool {
				return i%2 == 0
			},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Any(test.input, test.fn)
			assert.Equal(t, test.expected, out)
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    func() option.Option[int]
		fn       func(int) bool
		expected option.Option[int]
	}{
		{
			Label: testtypes.Label{
				Name: "empty iterator returns None",
			},
			input: From[int](),
			fn: func(i int) bool {
				return true
			},
			expected: option.None[int](),
		},
		{
			Label: testtypes.Label{
				Name: "finds the first match",
			},
			input: From(1, 2, 3, 4, 5),
			fn: func(i int) bool {
				return i%2 == 0
			},
			expected: option.Some(2),
		},
		{
			Label: testtypes.Label{
				Name: "no match returns none",
			},
			input: From(1, 2, 3),
			fn: func(i int) bool {
				return i > 5
			},
			expected: option.None[int](),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Find(test.input, test.fn)
			assert.Equal(t, test.expected, out)
		})
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    func() option.Option[int]
		n        uint
		expected func() option.Option[int]
	}{
		{
			Label: testtypes.Label{
				Name: "take zero returns an empty iterator",
			},
			input:    From(1, 2, 3),
			n:        0,
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "empty iterator returns an empty iterator",
			},
			input:    From[int](),
			n:        1,
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "taking more than iterator size returns all iterator's elements",
			},
			input:    From(1, 2, 3),
			n:        100,
			expected: From(1, 2, 3),
		},
		{
			Label: testtypes.Label{
				Name: "taking n elements in infinite iterator returns only the first n elements",
			},
			input: Unfold(0, func(i int) int {
				return i + 1
			}),
			n:        5,
			expected: From(0, 1, 2, 3, 4),
		},
		{
			Label: testtypes.Label{
				Name: "taking n elements in iterator returns only the first n elements",
			},
			input:    From(1, 2, 3, 4, 5),
			n:        2,
			expected: From(1, 2),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Take(test.input, test.n)
			assert.Equal(t, CollectToSlice(test.expected), CollectToSlice(out))
		})
	}
}

func TestTakeWhile(t *testing.T) {
	tests := []struct {
		testtypes.Label
		iter     func() option.Option[int]
		f        func(int) bool
		expected func() option.Option[int]
	}{
		{
			Label: testtypes.Label{
				Name: "take nothing",
			},
			iter:     From(1, 2, 3),
			f:        func(_ int) bool { return false },
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "take all",
			},
			iter:     From(1, 2, 3),
			f:        func(_ int) bool { return true },
			expected: From(1, 2, 3),
		},
		{
			Label: testtypes.Label{
				Name: "take some",
			},
			iter:     From(1, 3, 5, 7, 9, 2, 2, 1, 3, 5),
			f:        func(i int) bool { return i%2 != 0 },
			expected: From(1, 3, 5, 7, 9),
		},
	}
	for _, test := range tests {
		assert.Equal(
			t,
			CollectToSlice(test.expected),
			CollectToSlice(TakeWhile(test.iter, test.f)))
	}
}

func TestDrop(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    func() option.Option[int]
		n        uint
		expected func() option.Option[int]
	}{
		{
			Label: testtypes.Label{
				Name: "drop zero returns original iterator",
			},
			input:    From(1, 2, 3),
			n:        0,
			expected: From(1, 2, 3),
		},
		{
			Label: testtypes.Label{
				Name: "empty iterator returns an empty iterator",
			},
			input:    From[int](),
			n:        1,
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "dropping more than iterator size returns empty iterator's elements",
			},
			input:    From(1, 2, 3),
			n:        100,
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "dropping n elements in iterator skips the first n elements",
			},
			input:    From(1, 2, 3, 4, 5),
			n:        2,
			expected: From(3, 4, 5),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Drop(test.input, test.n)
			assert.Equal(t, CollectToSlice(test.expected), CollectToSlice(out))
		})
	}
}

func TestDropWhile(t *testing.T) {
	tests := []struct {
		testtypes.Label
		iter     func() option.Option[int]
		f        func(int) bool
		expected func() option.Option[int]
	}{
		{
			Label: testtypes.Label{
				Name: "drop nothing",
			},
			iter:     From(1, 2, 3),
			f:        func(_ int) bool { return false },
			expected: From(1, 2, 3),
		},
		{
			Label: testtypes.Label{
				Name: "drop all",
			},
			iter:     From(1, 2, 3),
			f:        func(_ int) bool { return true },
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "drop some",
			},
			iter:     From(1, 3, 5, 7, 9, 2, 2, 1, 3, 5),
			f:        func(i int) bool { return i%2 != 0 },
			expected: From(2, 2, 1, 3, 5),
		},
	}
	for _, test := range tests {
		assert.Equal(
			t,
			CollectToSlice(test.expected),
			CollectToSlice(DropWhile(test.iter, test.f)))
	}
}

func TestChain(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    []func() option.Option[int]
		expected func() option.Option[int]
	}{
		{
			Label: testtypes.Label{
				Name: "no iterators returns empty iterator",
			},
			input:    []func() option.Option[int]{},
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "empty iterators return empty iterator",
			},
			input: []func() option.Option[int]{
				From[int](),
				From[int](),
				From[int](),
			},
			expected: From[int](),
		},
		{
			Label: testtypes.Label{
				Name: "multiple iterators join correctly",
			},
			input: []func() option.Option[int]{
				From(1, 2, 3),
				From(4, 5),
				From(6, 7),
				From(8),
				From(9, 10),
			},
			expected: From(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
		},
		{
			Label: testtypes.Label{
				Name: "multiple iterators handle intermediate empty iterators",
			},
			input: []func() option.Option[int]{
				From[int](),
				From(1, 2, 3),
				From[int](),
				From(4, 5),
				From[int](),
			},
			expected: From(1, 2, 3, 4, 5),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Chain(test.input...)
			assert.Equal(t, CollectToSlice(test.expected), CollectToSlice(out))
		})
	}
}

func TestIsEqualBy(t *testing.T) {
	tests := []struct {
		testtypes.Label
		left     func() option.Option[int]
		right    func() option.Option[int]
		f        func(int, int) bool
		expected bool
	}{
		{
			Label: testtypes.Label{
				Name: "different values",
			},
			left:     From(1, 2, 3),
			right:    From(2, 3, 4),
			f:        func(i1, i2 int) bool { return i1 == i2 },
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "same values different order",
			},
			left:     From(1, 2, 3),
			right:    From(1, 3, 2),
			f:        func(i1, i2 int) bool { return i1 == i2 },
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "different length",
			},
			left:     From(1, 2, 3, 4, 5),
			right:    From(1, 2, 3),
			f:        func(i1, i2 int) bool { return i1 == i2 },
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "test equal iterators",
			},
			left:     From(1, 2, 3),
			right:    From(1, 2, 3),
			f:        func(i1, i2 int) bool { return i1 == i2 },
			expected: true,
		},
		{
			Label: testtypes.Label{
				Name: "test equal iterators with f",
			},
			left:     From(2, 4),
			right:    From(2, 4),
			f:        func(i1, i2 int) bool { return (i1+i2)%2 == 0 },
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.expected, IsEqualBy(test.left, test.right, test.f))
		})
	}
}

func TestIsEqual(t *testing.T) {
	tests := []struct {
		testtypes.Label
		left     func() option.Option[int]
		right    func() option.Option[int]
		expected bool
	}{
		{
			Label: testtypes.Label{
				Name: "different values",
			},
			left:     From(1, 2, 3),
			right:    From(2, 3, 4),
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "same values different order",
			},
			left:     From(1, 2, 3),
			right:    From(1, 3, 2),
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "different length",
			},
			left:     From(1, 2, 3, 4, 5),
			right:    From(1, 2, 3),
			expected: false,
		},
		{
			Label: testtypes.Label{
				Name: "test equal iterators",
			},
			left:     From(1, 2, 3),
			right:    From(1, 2, 3),
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.expected, IsEqual(test.left, test.right))
		})
	}
}
