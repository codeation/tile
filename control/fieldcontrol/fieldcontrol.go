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

type FieldControl struct {
	f *field.Field
}

func New(f *field.Field) *FieldControl {
	return &FieldControl{
		f: f,
	}
}

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
