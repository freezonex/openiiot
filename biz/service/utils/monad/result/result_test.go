package result

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"freezonex/openiiot/biz/service/utils/monad/option"
	"freezonex/openiiot/biz/service/utils/testtypes"

	"github.com/stretchr/testify/assert"
)

type testElement struct {
	value int
	ptr   *testElement
}

var (
	errorOne = errors.New("error one")
	errorTwo = errors.New("error two")
)

func TestOkWithPrimitive(t *testing.T) {
	value1 := "abc"
	r1 := Ok(value1)
	if !Equals(r1, Result[string]{value: option.Some(value1), err: nil}) {
		t.Error()
	}

	value2 := &value1
	r2 := Ok(value2)
	if !Equals(r2, Result[*string]{value: option.Some(value2), err: nil}) {
		t.Error()
	}
}

func TestOkWithSlice(t *testing.T) {
	value1 := []string{"one", "two"}
	r1 := Ok(value1)
	if !Equals(r1, Result[[]string]{value: option.Some(value1), err: nil}) {
		t.Error()
	}

	value2 := &value1
	r2 := Ok(value2)
	if !Equals(r2, Result[*[]string]{value: option.Some(value2), err: nil}) {
		t.Error()
	}
}

func TestOkWithMap(t *testing.T) {
	value1 := map[int]string{1: "one", 2: "two"}
	r1 := Ok(value1)
	if !Equals(r1, Result[map[int]string]{value: option.Some(value1), err: nil}) {
		t.Error()
	}

	value2 := &value1
	r2 := Ok(value2)
	if !Equals(r2, Result[*map[int]string]{value: option.Some(value2), err: nil}) {
		t.Error()
	}
}

func testOkWithStruct(t *testing.T) {
	value1 := testElement{value: 1, ptr: nil}
	o1 := Ok(value1)
	if !Equals(o1, Result[testElement]{value: option.Some(value1), err: nil}) {
		t.Error()
	}

	value2 := testElement{value: 1, ptr: nil}
	o2 := Ok(value2)
	if !Equals(o2, Result[testElement]{value: option.Some(value2), err: nil}) {
		t.Error()
	}
}

func TestErr(t *testing.T) {
	o := Err[int](errorOne)
	if !Equals(o, Result[int]{value: option.None[int](), err: errorOne}) {
		t.Error()
	}
}

func TestIsOk(t *testing.T) {
	if ok := Ok(1); !ok.IsOk() {
		t.Error()
	}
	if err := Err[int](errorOne); err.IsOk() {
		t.Error()
	}
}

func TestIsErr(t *testing.T) {
	if err := Err[int](errorOne); !err.IsErr() {
		t.Error()
	}
	if ok := Ok(1); ok.IsErr() {
		t.Error()
	}
}

func TestUnwrapValuePrimitive(t *testing.T) {
	value1 := "abc"
	r1 := Ok(value1)
	if r1.UnwrapValue() != value1 {
		t.Error()
	}

	value2 := &value1
	o2 := Ok(value2)
	if o2.UnwrapValue() != value2 {
		t.Error()
	}
}

func TestUnwrapSlice(t *testing.T) {
	value1 := []string{"one", "two"}
	r1 := Ok(value1)
	if !reflect.DeepEqual(r1.UnwrapValue(), value1) {
		t.Error()
	}

	value2 := &value1
	r2 := Ok(value2)
	if unwrapped := r2.UnwrapValue(); unwrapped != value2 || !reflect.DeepEqual(*unwrapped, *value2) {
		t.Error()
	}
}

func TestUnwrapMap(t *testing.T) {
	value1 := map[int]string{1: "one", 2: "two"}
	r1 := Ok(value1)
	if !reflect.DeepEqual(r1.UnwrapValue(), value1) {
		t.Error()
	}

	value2 := &value1
	r2 := Ok(value2)
	if unwrapped := r2.UnwrapValue(); unwrapped != value2 || !reflect.DeepEqual(*unwrapped, *value2) {
		t.Error()
	}
}

func TestUnwrapStruct(t *testing.T) {
	value1 := testElement{value: 1, ptr: &testElement{value: 2, ptr: nil}}
	r1 := Ok(value1)
	if !reflect.DeepEqual(r1.UnwrapValue(), value1) {
		t.Error()
	}

	value2 := &value1
	r2 := Ok(value2)
	if unwrapped := r2.UnwrapValue(); unwrapped != value2 || !reflect.DeepEqual(*unwrapped, *value2) {
		t.Error()
	}
}

func TestUnwrapError(t *testing.T) {
	r := Err[int](errorOne)
	if !errors.Is(r.UnwrapError(), errorOne) {
		t.Error()
	}
}

func TestFmapPrimitive(t *testing.T) {
	ok := Ok("abc")
	result := Fmap(ok, func(x string) int { return len(x) })
	if !Equals(result, Ok(3)) {
		t.Error()
	}
	err := Err[string](errorOne)
	result = Fmap(err, func(x string) int { return len(x) })
	if !Equals(result, Err[int](errorOne)) {
		t.Error()
	}
}

func TestFmapDifferentPointer(t *testing.T) {
	value := "abc"
	ptr := &value
	r1 := Ok(ptr)
	r2 := Fmap(r1, func(x *string) *string {
		y := strings.ToUpper(*x)
		return &y
	})
	if r2.UnwrapValue() == ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*r2.UnwrapValue(), "ABC") {
		t.Error()
	}
}

func TestFmapSamePointer(t *testing.T) {
	value := testElement{value: 1, ptr: nil}
	r1 := Ok(&value)
	r2 := Fmap(r1, func(x *testElement) *testElement {
		x.ptr = &testElement{value: 2, ptr: nil}
		return x
	})
	if r2.UnwrapValue() != &value {
		t.Error()
	}
	if !Equals(r2, Ok(&value)) {
		t.Error()
	}
}

func TestBindPrimitive(t *testing.T) {
	r := Bind(Ok("abc"), func(x string) Result[int] { return Ok(len(x)) })
	if !Equals(r, Ok(3)) {
		t.Error()
	}
	r = Bind(Ok("abc"), func(x string) Result[int] { return Err[int](errorOne) })
	if !Equals(r, Err[int](errorOne)) {
		t.Error()
	}
	r = Bind(Err[string](errorOne), func(x string) Result[int] { return Ok(len(x)) })
	if !Equals(r, Err[int](errorOne)) {
		t.Error()
	}
	r = Bind(Err[string](errorOne), func(x string) Result[int] { return Err[int](errorTwo) })
	if !Equals(r, Err[int](errorOne)) {
		t.Error()
	}
}

func TestBindSamePointer(t *testing.T) {
	value := testElement{value: 1, ptr: nil}
	ptr := &value
	r1 := Ok(ptr)
	r2 := Bind(r1, func(t *testElement) Result[*testElement] {
		t.value = 2
		return Ok(t)
	})
	if r2.UnwrapValue() != ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*r2.UnwrapValue(), testElement{value: 2, ptr: nil}) {
		t.Error()
	}
}

func TestBindDifferentPointer(t *testing.T) {
	value := "abc"
	ptr := &value
	r1 := Ok(ptr)
	r2 := Bind(r1, func(x *string) Result[*string] {
		y := strings.ToUpper(*x)
		return Ok(&y)
	})
	if r2.UnwrapValue() == ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*r2.UnwrapValue(), "ABC") {
		t.Error()
	}
}

func TestApplicativePrimitive(t *testing.T) {
	f := func(s string) int { return len(s) }
	result := Applicative(Err[func(string) int](errorOne), Err[string](errorTwo))
	if !Equals(result, Err[int](errorOne)) {
		t.Error()
	}
	result = Applicative(Err[func(string) int](errorOne), Ok("abc"))
	if !Equals(result, Err[int](errorOne)) {
		t.Error()
	}
	result = Applicative(Ok(f), Ok("abc"))
	if !Equals(result, Ok(3)) {
		t.Error()
	}
	result = Applicative(Ok(f), Err[string](errorOne))
	if !Equals(result, Err[int](errorOne)) {
		t.Error()
	}
}

func TestApplicativeSamePointer(t *testing.T) {
	f := func(s *testElement) *testElement {
		s.value = 2
		return s
	}
	value := testElement{value: 1, ptr: nil}
	ptr := &value
	r1 := Ok(f)
	r2 := Ok(ptr)
	r3 := Applicative(r1, r2)
	if r3.UnwrapValue() != ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*r3.UnwrapValue(), testElement{value: 2, ptr: nil}) {
		t.Error()
	}
}

func TestApplicativeDifferentPointer(t *testing.T) {
	f := func(x *string) *string {
		y := strings.ToUpper(*x)
		return &y
	}
	value := "abc"
	ptr := &value
	r1 := Ok(f)
	r2 := Ok(ptr)
	r3 := Applicative(r1, r2)
	if r3.UnwrapValue() == ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*r3.UnwrapValue(), "ABC") {
		t.Error()
	}
}

func TestEquals(t *testing.T) {
	if !Equals(Err[int](errorOne), Err[int](errorOne)) {
		t.Error()
	}
	if Equals(Err[int](errorOne), Err[int](errorTwo)) {
		t.Error()
	}
	if !Equals(Err[*int](errorOne), Err[*int](errorOne)) {
		t.Error()
	}
	if !Equals(Ok(1), Ok(1)) {
		t.Error()
	}
	if !Equals(Ok(&testElement{}), Ok(&testElement{})) {
		t.Error()
	}
	if Equals(Ok(1), Err[int](errorOne)) {
		t.Error()
	}
	if Equals(Err[int](errorOne), Ok(1)) {
		t.Error()
	}
}

func SequenceTest(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    []Result[int]
		expected Result[[]int]
	}{
		{
			Label: testtypes.Label{
				Name: "empty slice",
			},
			input:    []Result[int]{},
			expected: Ok([]int{}),
		},
		{
			Label: testtypes.Label{
				Name: "all something",
			},
			input: []Result[int]{
				Ok(1),
				Ok(2),
			},
			expected: Ok([]int{1, 2}),
		},
		{
			Label: testtypes.Label{
				Name: "has nothing",
			},
			input: []Result[int]{
				Ok(1),
				Err[int](errorOne),
				Ok(2),
				Err[int](errorTwo),
			},
			expected: Err[[]int](errorOne),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.expected, Sequence(test.input))
		})
	}
}
