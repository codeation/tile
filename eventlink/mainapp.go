package eventlink

import (
	"context"
	"image"
	"image/color"

	"github.com/codeation/impress"
	"github.com/codeation/impress/event"

	"github.com/codeation/tile/eventlink/ctxchan"
	"github.com/codeation/tile/eventlink/rectframe"
	"github.com/codeation/tile/eventlink/syncvar"
)

// RootApp is wrapper to impress.Application to implement AppFramer interface
type RootApp struct {
	application *impress.Application
	rect        *syncvar.Var[image.Point]
	cancelFunc  func()
}

// MainApp creates RootApp from impress.Application
func MainApp(a *impress.Application) *RootApp {
	return &RootApp{
		application: a,
		rect:        syncvar.New(image.Point{}),
		cancelFunc:  func() {},
	}
}

// Application returns impress.Application
func (app *RootApp) Application() *impress.Application {
	return app.application
}

// Cancel cancels child context.Context
func (app *RootApp) Cancel() {
	app.cancelFunc()
}

// Close closes MainApp resources include impress.Application
func (app *RootApp) Close() {
	app.application.Close()
}

// NewWindow creates child window
func (app *RootApp) NewWindow(rect image.Rectangle, background color.Color) *impress.Window {
	return app.application.NewWindow(rect, background)
}

// NewRectFrame creates child frame
func (app *RootApp) NewRectFrame(rect image.Rectangle) *rectframe.RectFrame {
	return rectframe.New(app.application, rect)
}

// Rects returns outer size of main frame
func (app *RootApp) Rect() image.Rectangle {
	return image.Rectangle{Max: app.rect.Get()}
}

// Rects returns inner size of  main frame
func (app *RootApp) InnerRect() image.Rectangle {
	return image.Rectangle{Max: app.rect.Get()}
}

// Run runs child actor
func (app *RootApp) Run(parentCtx context.Context, child Actor) {
	var ctx context.Context
	ctx, app.cancelFunc = context.WithCancel(parentCtx)
	link := New()
	link.Link(ctx, app, child)
	defer link.Close()

	for {
		e, ok := ctxchan.Get(ctx, app.application.Chan())
		if !ok {
			return
		}

		if e.Type() == event.ConfigureType {
			if ev, ok := e.(event.Configure); ok {
				app.rect.Set(ev.InnerSize)
			}
		}

		link.Put(ctx, e)
	}
}
