package iter

import "freezonex/openiiot/biz/service/utils/monad/option"

// FromSlice creates an iterator from a slice.
func FromSlice[A any](data []A) func() option.Option[A] {
	index := 0
	return func() option.Option[A] {
		if index >= 0 && index < len(data) {
			element := data[index]
			index++
			return option.Some(element)
		} else {
			return option.None[A]()
		}
	}
}

// CollectToSlice collects an iterator into a slice.
func CollectToSlice[A any](iter func() option.Option[A]) []A {
	out := make([]A, 0)
	elem := iter()
	for elem.IsSome() {
		out = append(out, elem.Unwrap())
		elem = iter()
	}
	return out
}
