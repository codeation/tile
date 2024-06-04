package eventlink

import (
	"context"

	"github.com/codeation/impress/event"
)

// PutInnerPt puts an event into a child event channel. Mouse event coordinates are shifted to match the child frame rectangle
func (c *EventLink) PutInnerPt(ctx context.Context, e event.Eventer) {
	switch ev := e.(type) {
	case event.Button:
		c.Put(ctx, event.Button{
			Action: ev.Action,
			Button: ev.Button,
			Point:  ev.Point.Sub(c.appFramer.Rect().Min),
		})
	case event.Motion:
		c.Put(ctx, event.Motion{
			Point:   ev.Point.Sub(c.appFramer.Rect().Min),
			Shift:   ev.Shift,
			Control: ev.Control,
			Alt:     ev.Alt,
			Meta:    ev.Meta,
		})
	default:
		c.Put(ctx, e)
	}
}
