package fieldview

import (
	"github.com/codeation/tile/view/fn"
)

// FieldWithFocused contains a field with focused state func.
type FieldWithFocused struct {
	Fielder
	isFocused fn.Bool
}

// WithFocused returns a FieldWithFocused.
func WithFocused(field Fielder, isFocused fn.Bool) *FieldWithFocused {
	return &FieldWithFocused{
		Fielder:   field,
		isFocused: isFocused,
	}
}

// Focused returns a focused flag
func (m *FieldWithFocused) Focused() bool { return m.isFocused() }

// FieldWithoutFocused contains a field with Focused always equal to False.
type FieldWithoutFocused struct {
	Fielder
}

// WithoutFocused returns a FieldWithoutFocused.
func WithoutFocused(field Fielder) *FieldWithFocused {
	return &FieldWithFocused{
		Fielder: field,
	}
}

// Focused returns a focused flag.
func (m *FieldWithoutFocused) Focused() bool { return false }
