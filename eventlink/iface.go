package eventlink

import (
	"context"
	"image"
	"image/color"

	"github.com/codeation/impress"
	"github.com/codeation/impress/event"

	"github.com/codeation/tile/eventlink/rectframe"
)

// App is an interface for accessing application methods, a parent event channel, and a parent frame
type App interface {
	Linker
	AppFramer
}

// Linker is an interface for accessing parent EventLink methods
type Linker interface {
	Chan() <-chan event.Eventer
	Put(ctx context.Context, ev event.Eventer)
	Link(ctx context.Context, actor Actor)
}

// AppFramer is an interface for accessing application methods and the parent frame
type AppFramer interface {
	Control
	Framer
}

// Control is an interface for accessing application methods
type Control interface {
	Sync()
	Cancel()
}

// Framer is an interface for accessing the parent frame
type Framer interface {
	NewWindow(rect image.Rectangle, background color.Color) *impress.Window
	NewRectFrame(rect image.Rectangle) *rectframe.RectFrame
	Rect() image.Rectangle
	InnerRect() image.Rectangle
}

// Actor is an interface for linking a child controller
type Actor interface {
	Action(ctx context.Context, app App)
	Wait()
}
