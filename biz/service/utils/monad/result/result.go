// result is a package providing the Result Monad.
package result

import (
	"errors"

	"freezonex/openiiot/biz/service/utils/monad/option"
)

// Result is a Monad to reflect success or failure/error, represented
// with the virtual union type of Ok(T) | Err(error)
type Result[A any] struct {
	value option.Option[A]
	err   error
}

// IsOk checks if the Result contains a successful value.
func (r Result[A]) IsOk() bool {
	return r.value.IsSome() && r.err == nil
}

// IsErr checks if the Result contains a failure error.
func (r Result[A]) IsErr() bool {
	return !r.IsOk()
}

// UnwrapValue unwraps the success value of the Option.
// Ok(T) -> T
//
// This is unsafe and checking using IsOk/IsErr
// is recommended before unwrapping.
func (r Result[A]) UnwrapValue() A {
	return r.value.Unwrap()
}

// UnwrapError unwraps the failure error of the Option.
// Err(error) -> error
//
// This is unsafe and checking using IsOk/IsErr
// is recommneded before unwrapping.
func (r Result[A]) UnwrapError() error {
	return r.err
}

// Ok is the constructor for a Result with a successful value.
// T -> Ok(T)
func Ok[A any](value A) Result[A] {
	return Result[A]{
		value: option.Some(value),
		err:   nil,
	}
}

// Err is the constructor for a Result with a failure error.
// error -> Err(error)
func Err[A any](err error) Result[A] {
	return Result[A]{
		value: option.None[A](),
		err:   err,
	}
}

// Fmap is the functor map / flat map operations on functors,
// applying an operation f on the value of a Result if it has a success value,
// returing a Result with the new value after f is applied, Ok(A) -> Ok(B), Err[A](e) -> Err[B](e).
// If the Result has an failure error, it is propogated.
func Fmap[A, B any](r Result[A], f func(A) B) Result[B] {
	if r.IsErr() {
		return Err[B](r.UnwrapError())
	}
	return Ok(f(r.UnwrapValue()))
}

// Bind is the Monadic bind that applys an operation f on the value of
// the Result if it has a success value, returning the value of f.
// If the Result has a failure error, it is propogated, Ok(A) -> Ok(B), Err[A](e) -> Err[B](e).
func Bind[A, B any](r Result[A], f func(A) Result[B]) Result[B] {
	if r.IsErr() {
		return Err[B](r.UnwrapError())
	}
	return f(r.UnwrapValue())
}

// Applicative is the applicative operator on a Result containing a computation f,
// and applying it to a Result with a success value that f takes as argument, returing a
// Result containing the output of f applied on the value if both results contains success values, Ok(f), Ok(A) -> Ok(f(a)).
// If either Result contains failure errors, the failure of the first Result r1 is propagated.
func Applicative[A any, B any](r1 Result[func(a A) B], r2 Result[A]) Result[B] {
	switch {
	case r1.IsErr():
		return Err[B](r1.UnwrapError())
	case r2.IsErr():
		return Err[B](r2.UnwrapError())
	default:
		return Ok(r1.UnwrapValue()(r2.UnwrapValue()))
	}
}

// Equals applies equality check on two Results.
func Equals[A any](r1, r2 Result[A]) bool {
	if r1.IsOk() && r2.IsOk() {
		return option.Equals(r1.value, r2.value)
	} else if r1.IsErr() && r2.IsErr() {
		return errors.Is(r1.err, r2.err)
	} else {
		return false
	}
}

// Sequence sequences a list of results, mapping the individual results
// into a Result of the elements.
func Sequence[A any](results []Result[A]) Result[[]A] {
	data := make([]A, 0)
	for _, result := range results {
		if result.IsErr() {
			return Err[[]A](result.UnwrapError())
		}
		data = append(data, result.UnwrapValue())
	}
	return Ok(data)
}
