// option is a package providing the Option Monad.
package option

import "reflect"

// Option is a Monad to reflect the possibility of a missing value,
// represented with the virtual union type of Some(T) | None.
type Option[A any] struct {
	value *A
}

// IsSome checks if the Option contains a value.
func (o Option[A]) IsSome() bool {
	return o.value != nil
}

// IsNone checks if the Option has nothing (equivalent of nil).
func (o Option[A]) IsNone() bool {
	return o.value == nil
}

// Unwrap unwraps the value of the Option.
// Some(T) -> T
//
// This is unsafe and checking using IsSome/IsNone
// is recommended before unwrapping.
func (o Option[A]) Unwrap() A {
	return *o.value
}

// Some is the constructor for an Option with a value.
func Some[A any](value A) Option[A] {
	return Option[A]{value: &value}
}

// None is the constructor for an Option with no value.
func None[A any]() Option[A] {
	return Option[A]{value: nil}
}

// Fmap is the functor map / flat map operations on functors,
// applying an operation f on the value of an Option if it is present,
// returing an Option with the new value after f is applied, Some(A) -> Some(B), None[A]() -> None[B]().
// If the Option has no value, it is propagated.
func Fmap[A, B any](o Option[A], f func(A) B) Option[B] {
	if o.IsNone() {
		return None[B]()
	}
	return Some(f(o.Unwrap()))
}

// Bind is the Monadic bind that applys an operation f on the value of
// the Option if it is present, returning the value of f.
// If the Option has nothing, it is propagated, Some(A) -> Some(B), None[A]() -> None[B]().
func Bind[A, B any](o Option[A], f func(A) Option[B]) Option[B] {
	if o.IsNone() {
		return None[B]()
	}
	return f(o.Unwrap())
}

// Applicative is the applicative operator on an Option containing a computation f,
// and applying it to an Option with a value that f takes as argument, returing an
// Option containing the output of f applied on the value if both options contains values, Some(f), Some(a) -> Some(f(a)).
// If either Option contains nothing, nothing is propagated.
func Applicative[A, B any](o1 Option[func(a A) B], o2 Option[A]) Option[B] {
	if o1.IsNone() || o2.IsNone() {
		return None[B]()
	}
	return Some(o1.Unwrap()(o2.Unwrap()))
}

// Equals applies equality check on two Options.
func Equals[A any](o1, o2 Option[A]) bool {
	switch {
	case o1.IsNone() && o2.IsNone():
		return true
	case o1.IsSome() && o2.IsSome():
		return reflect.DeepEqual(o1.Unwrap(), o2.Unwrap())
	default:
		return false
	}
}

// Sequence sequences a slice of options, lifting the slice of options
// into an Option of the elements.
func Sequence[A any](options []Option[A]) Option[[]A] {
	data := make([]A, 0)
	for _, option := range options {
		if option.IsNone() {
			return None[[]A]()
		}
		data = append(data, option.Unwrap())
	}
	return Some(data)
}
