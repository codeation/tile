package fieldview

import (
	"github.com/codeation/tile/view/fn"
)

// FieldWithFocused contains a field with focused state func
type FieldWithFocused struct {
	Fielder
	isFocused fn.Bool
}

// WithFocused returns a FieldWithFocused
func WithFocused(field Fielder, isFocused fn.Bool) *FieldWithFocused {
	return &FieldWithFocused{
		Fielder:   field,
		isFocused: isFocused,
	}
}

// Focused returns a focused flag
func (m *FieldWithFocused) Focused() bool { return m.isFocused() }
