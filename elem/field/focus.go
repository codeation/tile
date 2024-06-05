package field

type FocusState struct {
	focused bool
}

func (m *FocusState) Focus()        { m.focused = true }
func (m *FocusState) Blur()         { m.focused = false }
func (m *FocusState) Focused() bool { return m.focused }

type FieldWithFocus struct {
	*Field
	*FocusState
}

func WithFocus(field *Field) *FieldWithFocus {
	return &FieldWithFocus{
		Field:      field,
		FocusState: new(FocusState),
	}
}
