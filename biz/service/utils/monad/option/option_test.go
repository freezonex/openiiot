package option

import (
	"reflect"
	"strings"
	"testing"

	"freezonex/openiiot/biz/service/utils/testtypes"

	"github.com/stretchr/testify/assert"
)

type testElement struct {
	value int
	ptr   *testElement
}

func TestSomeWithPrimitive(t *testing.T) {
	value1 := "abc"
	o1 := Some(value1)
	if !reflect.DeepEqual(o1, Option[string]{value: &value1}) {
		t.Error()
	}

	value2 := &value1
	o2 := Some(value2)
	if !reflect.DeepEqual(o2, Option[*string]{value: &value2}) {
		t.Error()
	}
}

func TestSomeWithSlice(t *testing.T) {
	value1 := []string{"one", "two"}
	o1 := Some(value1)
	if !reflect.DeepEqual(o1, Option[[]string]{value: &value1}) {
		t.Error()
	}

	value2 := &value1
	o2 := Some(value2)
	if !reflect.DeepEqual(o2, Option[*[]string]{value: &value2}) {
		t.Error()
	}
}

func TestSomeWithMap(t *testing.T) {
	value1 := map[int]string{1: "one", 2: "two"}
	o1 := Some(value1)
	if !reflect.DeepEqual(o1, Option[map[int]string]{value: &value1}) {
		t.Error()
	}

	value2 := &value1
	o2 := Some(value2)
	if !reflect.DeepEqual(o2, Option[*map[int]string]{value: &value2}) {
		t.Error()
	}
}

func testSomeWithStruct(t *testing.T) {
	value1 := testElement{value: 1, ptr: nil}
	o1 := Some(value1)
	if !reflect.DeepEqual(o1, Option[testElement]{value: &value1}) {
		t.Error()
	}

	value2 := testElement{value: 1, ptr: nil}
	o2 := Some(value2)
	if !reflect.DeepEqual(o2, Option[testElement]{value: &value2}) {
		t.Error()
	}
}

func TestNone(t *testing.T) {
	o := None[int]()
	if !reflect.DeepEqual(o, Option[int]{value: nil}) {
		t.Error()
	}
}

func TestIsSome(t *testing.T) {
	if some := Some(1); !some.IsSome() {
		t.Error()
	}
	if none := None[int](); none.IsSome() {
		t.Error()
	}
}

func TestIsNone(t *testing.T) {
	if none := None[int](); !none.IsNone() {
		t.Error()
	}
	if some := Some(1); some.IsNone() {
		t.Error()
	}
}

func TestUnwrapPrimitive(t *testing.T) {
	value1 := "abc"
	o1 := Some(value1)
	if o1.Unwrap() != value1 {
		t.Error()
	}

	value2 := &value1
	o2 := Some(value2)
	if o2.Unwrap() != value2 {
		t.Error()
	}
}

func TestUnwrapSlice(t *testing.T) {
	value1 := []string{"one", "two"}
	o1 := Some(value1)
	if !reflect.DeepEqual(o1.Unwrap(), value1) {
		t.Error()
	}

	value2 := &value1
	o2 := Some(value2)
	if unwrapped := o2.Unwrap(); unwrapped != value2 || !reflect.DeepEqual(*unwrapped, *value2) {
		t.Error()
	}
}

func TestUnwrapMap(t *testing.T) {
	value1 := map[int]string{1: "one", 2: "two"}
	o1 := Some(value1)
	if !reflect.DeepEqual(o1.Unwrap(), value1) {
		t.Error()
	}

	value2 := &value1
	o2 := Some(value2)
	if unwrapped := o2.Unwrap(); unwrapped != value2 || !reflect.DeepEqual(*unwrapped, *value2) {
		t.Error()
	}
}

func TestUnwrapStruct(t *testing.T) {
	value1 := testElement{value: 1, ptr: &testElement{value: 2, ptr: nil}}
	o1 := Some(value1)
	if !reflect.DeepEqual(o1.Unwrap(), value1) {
		t.Error()
	}

	value2 := &value1
	o2 := Some(value2)
	if unwrapped := o2.Unwrap(); unwrapped != value2 || !reflect.DeepEqual(*unwrapped, *value2) {
		t.Error()
	}
}

func TestFmapPrimitive(t *testing.T) {
	some := Some("abc")
	result := Fmap(some, func(x string) int { return len(x) })
	if !reflect.DeepEqual(result, Some(3)) {
		t.Error()
	}
	none := None[string]()
	result = Fmap(none, func(x string) int { return len(x) })
	if !reflect.DeepEqual(result, None[int]()) {
		t.Error()
	}
}

func TestFmapDifferentPointer(t *testing.T) {
	value := "abc"
	ptr := &value
	o1 := Some(ptr)
	o2 := Fmap(o1, func(x *string) *string {
		y := strings.ToUpper(*x)
		return &y
	})
	if o2.Unwrap() == ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*o2.Unwrap(), "ABC") {
		t.Error()
	}
}

func TestFmapSamePointer(t *testing.T) {
	value := testElement{value: 1, ptr: nil}
	o1 := Some(&value)
	o2 := Fmap(o1, func(x *testElement) *testElement {
		x.ptr = &testElement{value: 2, ptr: nil}
		return x
	})
	if o2.Unwrap() != &value {
		t.Error()
	}
	if !reflect.DeepEqual(o2, Some(&value)) {
		t.Error()
	}
}

func TestBindPrimitive(t *testing.T) {
	result := Bind(Some("abc"), func(x string) Option[int] { return Some(len(x)) })
	if !reflect.DeepEqual(result, Some(3)) {
		t.Error()
	}
	result = Bind(Some("abc"), func(x string) Option[int] { return None[int]() })
	if !reflect.DeepEqual(result, None[int]()) {
		t.Error()
	}
	result = Bind(None[string](), func(x string) Option[int] { return Some(len(x)) })
	if !reflect.DeepEqual(result, None[int]()) {
		t.Error()
	}
	result = Bind(None[string](), func(x string) Option[int] { return None[int]() })
	if !reflect.DeepEqual(result, None[int]()) {
		t.Error()
	}
}

func TestBindSamePointer(t *testing.T) {
	value := testElement{value: 1, ptr: nil}
	ptr := &value
	o1 := Some(ptr)
	o2 := Bind(o1, func(t *testElement) Option[*testElement] {
		t.value = 2
		return Some(t)
	})
	if o2.Unwrap() != ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*o2.Unwrap(), testElement{value: 2, ptr: nil}) {
		t.Error()
	}
}

func TestBindDifferentPointer(t *testing.T) {
	value := "abc"
	ptr := &value
	o1 := Some(ptr)
	o2 := Bind(o1, func(x *string) Option[*string] {
		y := strings.ToUpper(*x)
		return Some(&y)
	})
	if o2.Unwrap() == ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*o2.Unwrap(), "ABC") {
		t.Error()
	}
}

func TestApplicativePrimitive(t *testing.T) {
	f := func(s string) int { return len(s) }
	result := Applicative(None[func(string) int](), None[string]())
	if !reflect.DeepEqual(result, None[int]()) {
		t.Error()
	}
	result = Applicative(None[func(string) int](), Some("abc"))
	if !reflect.DeepEqual(result, None[int]()) {
		t.Error()
	}
	result = Applicative(Some(f), Some("abc"))
	if !reflect.DeepEqual(result, Some(3)) {
		t.Error()
	}
	result = Applicative(Some(f), None[string]())
	if !reflect.DeepEqual(result, None[int]()) {
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
	o1 := Some(f)
	o2 := Some(ptr)
	o3 := Applicative(o1, o2)
	if o3.Unwrap() != ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*o3.Unwrap(), testElement{value: 2, ptr: nil}) {
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
	o1 := Some(f)
	o2 := Some(ptr)
	o3 := Applicative(o1, o2)
	if o3.Unwrap() == ptr {
		t.Error()
	}
	if !reflect.DeepEqual(*o3.Unwrap(), "ABC") {
		t.Error()
	}
}

func TestEquals(t *testing.T) {
	if !Equals(None[int](), None[int]()) {
		t.Error()
	}
	if !Equals(None[*int](), None[*int]()) {
		t.Error()
	}
	if !Equals(Some(1), Some(1)) {
		t.Error()
	}
	if !Equals(Some(&testElement{}), Some(&testElement{})) {
		t.Error()
	}
	if Equals(Some(1), None[int]()) {
		t.Error()
	}
	if Equals(None[int](), Some(1)) {
		t.Error()
	}
}

func TestSequence(t *testing.T) {
	tests := []struct {
		testtypes.Label
		input    []Option[int]
		expected Option[[]int]
	}{
		{
			Label: testtypes.Label{
				Name: "empty slice",
			},
			input:    []Option[int]{},
			expected: Some([]int{}),
		},
		{
			Label: testtypes.Label{
				Name: "all something",
			},
			input: []Option[int]{
				Some(1),
				Some(2),
			},
			expected: Some([]int{1, 2}),
		},
		{
			Label: testtypes.Label{
				Name: "has nothing",
			},
			input: []Option[int]{
				Some(1),
				None[int](),
				Some(2),
			},
			expected: None[[]int](),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.expected, Sequence(test.input))
		})
	}
}
