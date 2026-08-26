package dialog

import (
	"context"
	"image"

	"github.com/codeation/impress/event"
	"github.com/codeation/tile/eventlink"
	"github.com/codeation/tile/eventlink/ctxchan"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/area"
)

// Control contains dialog window control.
type Control struct {
	*Dialog
	w    area.Area
	View view.Viewer
	rect image.Rectangle
}

// NewControl creates a new dialog window control.
func NewControl(w area.Area, d *Dialog, viewer view.Viewer, rect image.Rectangle) *Control {
	return &Control{
		Dialog: d,
		w:      w,
		View:   viewer,
		rect:   rect,
	}
}

// Action is a loop for managing dialog window.
func (c *Control) Action(ctx context.Context, app eventlink.App) {
	for {
		if len(app.Chan()) == 0 {
			c.w.Window().Clear()
			c.View.Draw(c.w.Window(), image.Point{})
			c.w.Window().Show()
			app.Application().Sync()
		}

		e, ok := ctxchan.Get(ctx, app.Chan())
		if !ok {
			return
		}

		c.Control(ctx, app, e, c.Do)
	}
}

// Wait is a shutdown func.
func (c *Control) Wait() {
	c.w.Drop()
}

// Do implements a reaction to one event.
func (c *Control) Do(ctx context.Context, app eventlink.App, e event.Eventer) {
	switch e {
	case event.KeyEscape:
		app.Cancel()
	case event.KeyEnter:
		if c.IsLastFieldSelected() || c.IsAnyButtonSelected() {
			c.ButtonSelected().Call()
			return
		}
		next := c.Next(c.Selected())
		if next != nil {
			c.Select(next)
		}
	default:
		if ev, ok := e.(event.Button); ok {
			if ev.Action == event.ButtonActionPress && ev.Button == event.ButtonLeft {
				if elem, _ := c.View.Select(ev.Point, c.rect.Min); elem != nil {
					if c.TrySelect(elem) {
						if c.IsAnyButtonSelected() {
							c.ButtonSelected().Call()
						}
					}
				}
			}
		}
	}
}
