// pair package offers a way to represent tuples of arbitrary arity.
package pair

// Pair represents a pair of any two types.
type Pair[A, B any] struct {
	// First is the left value of the pair.
	First A
	// Second is the right value of the pair.
	Second B
}

// New generates a new Pair of elements of any types.
func New[A, B any](first A, second B) Pair[A, B] {
	return Pair[A, B]{
		First:  first,
		Second: second,
	}
}

// GetFirst gets the first value in a pair.
func GetFirst[A, B any](p Pair[A, B]) A {
	return p.First
}

// GetSecond gets the second value in a pair.
func GetSecond[A, B any](p Pair[A, B]) B {
	return p.Second
}
