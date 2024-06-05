package fn

// Const returns a function with a constant value of any type
func Const[T any](v T) func() T { return func() T { return v } }

func If[T any](condFn Bool, left, right func() T) func() T {
	return func() T {
		if condFn() {
			return left()
		}
		return right()
	}
}
