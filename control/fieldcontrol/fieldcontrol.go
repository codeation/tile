package fieldcontrol

import (
	"context"
	"unicode"

	"github.com/codeation/impress/clipboard"
	"github.com/codeation/impress/event"
	"github.com/codeation/tile/control"
	"github.com/codeation/tile/control/key"
	"github.com/codeation/tile/elem/field"
	"github.com/codeation/tile/eventlink"
)

// FieldControl provides control logic for a Field element, handling various events such as keyboard inputs and clipboard actions.
type FieldControl struct {
	f *field.Field
}

// New creates and returns a new FieldControl for the given Field element.
func New(f *field.Field) *FieldControl {
	return &FieldControl{
		f: f,
	}
}

// Control processes events for the FieldControl, updating the Field element based on keyboard and clipboard inputs.
// It delegates to the prior control function for unhandled events.
func (c *FieldControl) Control(ctx context.Context, app eventlink.App, e event.Eventer, prior control.DoFunc) {
	switch ev := e.(type) {
	case event.Keyboard:
		switch {
		case ev == event.KeyBackSpace:
			c.f.Backspace()
		case ev == event.KeyLeft:
			c.f.Left()
		case ev == event.KeyRight:
			c.f.Right()
		case ev == key.ShiftEnter:
			c.f.InsertNL()
		case ev.IsGraphic():
			c.f.Insert(ev.Rune)
		default:
			prior(ctx, app, e)
		}
	case event.Clipboard:
		if text, ok := ev.Data.(clipboard.Text); ok {
			for _, r := range text {
				if !unicode.IsGraphic(r) {
					continue
				}
				c.f.Insert(r)
			}
			return
		}
		prior(ctx, app, e)
	default:
		prior(ctx, app, e)
	}
}
