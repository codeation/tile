package fn

// Const returns a function with a constant value of any type.
func Const[T any](v T) func() T { return func() T { return v } }

// If creates a func that returns one of its parameters depending on result of condFn func.
func If[T any](condFn Bool, left, right T) func() T {
	return func() T {
		if condFn() {
			return left
		}
		return right
	}
}
