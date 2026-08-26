// Package contains a control for operating the dialog window
package dialog

import (
	"github.com/codeation/tile/control/fieldcontrol"
	"github.com/codeation/tile/control/sequence"
	"github.com/codeation/tile/elem/field"
)

// Field is a struct to add edit element to dialog window.
type Field struct {
	*fieldcontrol.FieldControl
	Prompt string
	Field  *field.Field
}

// NewField returns a new edit element.
func NewField(prompt string, f *field.Field) *Field {
	return &Field{
		FieldControl: fieldcontrol.New(f),
		Prompt:       prompt,
		Field:        f,
	}
}

// Button is a struct to add button to dialog window.
type Button struct {
	Prompt string
	Call   func()
}

// NewButton returns a new button.
func NewButton(prompt string, callFunc func()) *Button {
	return &Button{
		Prompt: prompt,
		Call:   callFunc,
	}
}

// Dialog contains fields and buttons, also selected elements.
type Dialog struct {
	*sequence.Sequence
	lastField *Field
	buttonSeq *sequence.Sequence
}

// New creates a model of dialog, struct is used to create a window control and view.
func New(fields []*Field, buttons []*Button) *Dialog {
	seq := sequence.New().Vertical()
	for _, f := range fields {
		seq.Add(f)
	}
	buttonSeq := sequence.New().Horizontal()
	for _, b := range buttons {
		buttonSeq.Add(b)
	}
	seq.Add(buttonSeq)
	return &Dialog{
		Sequence:  seq,
		lastField: fields[len(fields)-1],
		buttonSeq: buttonSeq,
	}
}

// IsLastFieldSelected returns a true, when latest edit element is active.
func (d *Dialog) IsLastFieldSelected() bool {
	return d.Selected() == d.lastField
}

// IsAnyButtonSelected returns a true, when buttons line is active.
func (d *Dialog) IsAnyButtonSelected() bool {
	return d.Selected() == d.buttonSeq
}

// ButtonSelected returns a current button, even if row with buttons is not active.
func (d *Dialog) ButtonSelected() *Button {
	return d.buttonSeq.Selected().(*Button)
}

// TrySelect selects a field or button as the active element of a dialog box.
func (d *Dialog) TrySelect(elem any) bool {
	for e := d.buttonSeq.Front(); e != nil; e = d.buttonSeq.Next(e) {
		if e == elem {
			d.buttonSeq.Select(e)
			d.Select(d.buttonSeq)
			return true
		}
	}
	for e := d.Front(); e != nil; e = d.Next(e) {
		if e == d.buttonSeq {
			continue
		}
		if e == elem {
			d.Select(e)
			return true
		}
	}
	return false
}
