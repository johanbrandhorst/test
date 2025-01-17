package interfaces

import (
	"math"

	"github.com/shoenig/test/internal/constraints"
)

// EqualFunc represents a type implementing the Equal method.
type EqualFunc[A any] interface {
	Equal(A) bool
}

// LessFunc represents any type implementing the Less method.
type LessFunc[A any] interface {
	Less(A) bool
}

// Map represents any map type where keys are comparable.
type Map[K comparable, V any] interface {
	~map[K]V
}

// MapEqualFunc represents any map type where keys are comparable and values implement .Equal method.
type MapEqualFunc[K comparable, V EqualFunc[V]] interface {
	~map[K]V
}

// Number is float, integer, or complex.
type Number interface {
	constraints.Ordered
	constraints.Float | constraints.Integer | constraints.Complex
}

// Numeric returns false if n is Inf/NaN.
//
// Always returns true for integral values.
func Numeric[N Number](n N) bool {
	check := func(f float64) bool {
		if math.IsNaN(f) {
			return false
		} else if math.IsInf(f, 0) {
			return false
		}
		return true
	}
	return check(float64(n))
}

// The LengthFunc interface is satisfied by a type that implements Len().
type LengthFunc interface {
	Len() int
}

// The SizeFunc interface is satisfied by a type that implements Size().
type SizeFunc interface {
	Size() int
}

// The EmptyFunc interface is satisfied by a type that implements Empty().
type EmptyFunc interface {
	Empty() bool
}

// The ContainsFunc interface is satisfied by a type that implements Contains(T).
type ContainsFunc[T any] interface {
	Contains(T) bool
}
