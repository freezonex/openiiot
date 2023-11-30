// iter is a package for performing useful operations on iterators
// and providing a way to transform iterators to and from
// data structures (slices/maps).
package iter

import (
	"errors"
	"sort"

	o "freezonex/openiiot/biz/service/utils/monad/option"
	"freezonex/openiiot/biz/service/utils/pair"
)

var ErrEmptyIterator = errors.New("empty iterator")

// From creates an iterator over the supplied arguments in order.
func From[A any](data ...A) func() o.Option[A] {
	return FromSlice(data)
}

// ForEach applys a function f over an iterable.
func ForEach[A any](iter func() o.Option[A], f func(A)) {
	elem := iter()
	for elem.IsSome() {
		f(elem.Unwrap())
		elem = iter()
	}
}

// Enumerate returns the provided iterator paired with a 0 based incrementing index.
func Enumerate[A any](iter func() o.Option[A]) func() o.Option[pair.Pair[uint, A]] {
	indexes := Unfold(0, func(i uint) uint { return i + 1 })
	return Zip(indexes, iter)
}

// Map maps a function over an iterator producing a type.
func Map[A, B any](iter func() o.Option[A], f func(A) B) func() o.Option[B] {
	return func() o.Option[B] {
		return o.Fmap(iter(), f)
	}
}

// Filter filters an iterable based on predicate on an iterator producing a type.
func Filter[A any](iter func() o.Option[A], f func(A) bool) func() o.Option[A] {
	return func() o.Option[A] {
		elem := iter()
		for elem.IsSome() && !f(elem.Unwrap()) {
			elem = iter()
		}
		return elem
	}
}

// Fold right folds on an iterator into a initial value, applying f at each fold.
func Fold[A, B any](iter func() o.Option[A], initial B, f func(B, A) B) B {
	acc := initial
	elem := iter()
	for elem.IsSome() {
		acc = f(acc, elem.Unwrap())
		elem = iter()
	}
	return acc
}

// Reduce right folds an iterator without an initial value, applying f at each fold.
//
// Returns an option value as value may not always exist.
func Reduce[A any](iter func() o.Option[A], f func(A, A) A) o.Option[A] {
	return o.Fmap(iter(), func(value A) A {
		return Fold(iter, value, f)
	})
}

// Unfold unfolds an iterator from an initial seed value and subsequently
// yields a new element from the iterator by applying a function over
// the previous value.
func Unfold[A any](seed A, generator func(A) A) func() o.Option[A] {
	current := seed
	return func() o.Option[A] {
		result := current
		current = generator(result)
		return o.Some(result)
	}
}

// Zip zips two iterators into a pair iterator.
func Zip[A, B any](left func() o.Option[A], right func() o.Option[B]) func() o.Option[pair.Pair[A, B]] {
	return func() o.Option[pair.Pair[A, B]] {
		leftElem := left()
		rightElem := right()
		if leftElem.IsNone() || rightElem.IsNone() {
			return o.None[pair.Pair[A, B]]()
		} else {
			return o.Some(pair.New(leftElem.Unwrap(), rightElem.Unwrap()))
		}
	}
}

// All checks all elements in an iterator satisfies a condition.
// Empty iterator always returns true.
func All[A any](iter func() o.Option[A], f func(A) bool) bool {
	return Fold(iter, true, func(acc bool, elem A) bool {
		return acc && f(elem)
	})
}

// Any checks if at least one element in an iterator satisfies a condition.
// Empty iterator always return false.
func Any[A any](iter func() o.Option[A], f func(A) bool) bool {
	return Find(iter, f).IsSome()
}

// Find finds the first match.
func Find[A any](iter func() o.Option[A], f func(A) bool) o.Option[A] {
	elem := iter()
	for elem.IsSome() {
		if f(elem.Unwrap()) {
			return elem
		}
		elem = iter()
	}
	return o.None[A]()
}

// Take takes only the first n elements in iterator, which is particularly
// useful for infinite iterators.
func Take[A any](iter func() o.Option[A], n uint) func() o.Option[A] {
	count := uint(0)
	return func() o.Option[A] {
		if count < n {
			element := iter()
			count++
			return element
		} else {
			return o.None[A]()
		}
	}
}

// Takes an iterator while a condition is met, whereby a failure on condition will terminate
// the rest of the iterator.
func TakeWhile[A any](iter func() o.Option[A], f func(A) bool) func() o.Option[A] {
	hasTerminated := false
	return func() o.Option[A] {
		if hasTerminated {
			return o.None[A]()
		} else {
			elem := iter()
			if elem.IsNone() || !f(elem.Unwrap()) {
				hasTerminated = true
				return o.None[A]()
			} else {
				return elem
			}
		}
	}
}

// Drop drops/skips the first n elements in iterator, resulting in the
// first yielded element of the iterator to be the nth (zero-based)
// element in the original iterator.
func Drop[A any](iter func() o.Option[A], n uint) func() o.Option[A] {
	for count := uint(0); count < n; count++ {
		_ = iter()
	}
	return iter
}

// DropWhile will keep dropping elements in an iterator while a condition is met, whereby
// a failure will start yielding elements in the iterator from that point onwards.
func DropWhile[A any](iter func() o.Option[A], f func(A) bool) func() o.Option[A] {
	shouldKeepDropping := true
	return func() o.Option[A] {
		if !shouldKeepDropping {
			return iter()
		} else {
			elem := iter()
			for elem.IsSome() && f(elem.Unwrap()) {
				elem = iter()
			}
			shouldKeepDropping = false
			return elem
		}
	}
}

// Chain chains n iterators together, with the iterators being consumed in order.
func Chain[A any](iters ...func() o.Option[A]) func() o.Option[A] {
	index := 0
	return func() o.Option[A] {
		if index >= len(iters) {
			return o.None[A]()
		} else {
			for index < len(iters) {
				iter := iters[index]
				elem := iter()
				if elem.IsNone() {
					index++
					continue
				} else {
					return elem
				}
			}
			return o.None[A]()
		}
	}
}

// SortBy sorts an iterator from a given comparator function.
// Comparator function returns if a is less than b.
func SortBy[A any](iter func() o.Option[A], comparator func(a, b A) bool) func() o.Option[A] {
	data := CollectToSlice(iter)
	sort.SliceStable(data, func(i, j int) bool {
		return comparator(data[i], data[j])
	})
	return From(data...)
}

// IsEqualBy checks if two iterators are equal by applying a function that
// evaluates equality of each element of the two iterators for all elements in order.
func IsEqualBy[A any](left, right func() o.Option[A], checkingFunc func(A, A) bool) bool {
	leftElem := left()
	rightElem := right()
	for leftElem.IsSome() && rightElem.IsSome() {
		if !checkingFunc(leftElem.Unwrap(), rightElem.Unwrap()) {
			return false
		}
		leftElem = left()
		rightElem = right()
	}
	return leftElem.IsNone() && rightElem.IsNone()
}

// IsEqual checks if two iterators are equal for all elements in order.
func IsEqual[A comparable](left, right func() o.Option[A]) bool {
	return IsEqualBy(left, right, func(x, y A) bool {
		return x == y
	})
}
