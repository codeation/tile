package fieldview

// Fielder is an interface of a field model
type Fielder interface {
	Strings() []string
	Cursor() int
}

// FocusFielder is an interface of a field model with a focused flag
type FocusFielder interface {
	Fielder
	Focused() bool
}
