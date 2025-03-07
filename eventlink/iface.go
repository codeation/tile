package eventlink

import (
	"context"
	"image"
	"image/color"

	"github.com/codeation/impress"
	"github.com/codeation/impress/event"

	"github.com/codeation/tile/eventlink/rectframe"
)

// App is an interface that combines Linker and AppFramer interfaces,
// providing methods for accessing application methods, a parent event channel, and a parent frame.
type App interface {
	Linker
	AppFramer
}

// AppFramer is an interface that combines Control and Framer interfaces,
// providing methods for accessing application methods and the parent frame.
type AppFramer interface {
	Control
	Framer
}

// Linker is an interface that provides methods for interacting with parent EventLink methods.
type Linker interface {
	// Chan returns a channel for receiving events from the parent event link.
	Chan() <-chan event.Eventer
	// Put sends an event to the child event link if the context is not canceled.
	Put(ctx context.Context, ev event.Eventer)
	// Link links the parent event link with a new child controller.
	Link(ctx context.Context, appFramer AppFramer, actor Actor)
}

// Control is an interface that provides methods for accessing application-level methods.
type Control interface {
	// Application returns the current application instance.
	Application() *impress.Application
	// Cancel cancels a child controller context.
	Cancel()
}

// Framer is an interface that provides methods for accessing and manipulating the parent frame.
type Framer interface {
	// NewWindow creates a new window with the specified rectangle and background color.
	NewWindow(rect image.Rectangle, background color.Color) *impress.Window
	// NewRectFrame creates a new rectangular frame with the specified rectangle.
	NewRectFrame(rect image.Rectangle) *rectframe.RectFrame
	// Rect returns the rectangle representing the frame boundaries.
	Rect() image.Rectangle
	// InnerRect returns the inner rectangle of the frame.
	InnerRect() image.Rectangle
}

// Actor is an interface for linking and managing a child controller.
type Actor interface {
	// Action performs the main action for the child controller using the provided context and app.
	Action(ctx context.Context, app App)
	// Wait waits for the child controller's action to complete.
	Wait()
}
