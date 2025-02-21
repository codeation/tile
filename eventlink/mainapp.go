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

type mainApp struct {
	application *impress.Application
	rect        *syncvar.Var[image.Point]
	canceFunc   func()
}

// MainApp creates root AppFramer from impress.Application
func MainApp(a *impress.Application) *mainApp {
	return &mainApp{
		application: a,
		rect:        syncvar.New(image.Point{}),
		canceFunc:   func() {},
	}
}

func (app *mainApp) Application() *impress.Application {
	return app.application
}

func (app *mainApp) Cancel() {
	app.canceFunc()
}

func (app *mainApp) Close() {
	app.application.Close()
}

func (app *mainApp) NewWindow(rect image.Rectangle, background color.Color) *impress.Window {
	return app.application.NewWindow(rect, background)
}

func (app *mainApp) NewRectFrame(rect image.Rectangle) *rectframe.RectFrame {
	return rectframe.New(app.application, rect)
}

func (app *mainApp) Rect() image.Rectangle {
	return image.Rectangle{Max: app.rect.Get()}
}

func (app *mainApp) InnerRect() image.Rectangle {
	return image.Rectangle{Max: app.rect.Get()}
}

func (app *mainApp) Run(ctx context.Context, child Actor) {
	ctx, app.canceFunc = context.WithCancel(ctx)
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
