package eventlink

import (
	"context"
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/impress/event"

	"github.com/codeation/tile/eventlink/ctxchan"
	"github.com/codeation/tile/eventlink/rectframe"
	"github.com/codeation/tile/eventlink/syncvar"
)

type mainApp struct {
	*impress.Application
	rect      *syncvar.Var[image.Point]
	canceFunc func()
}

// MainApp creates root AppFramer from impress.Application
func MainApp(a *impress.Application) *mainApp {
	return &mainApp{
		Application: a,
		rect:        syncvar.New(image.Point{}),
		canceFunc:   func() {},
	}
}

func (app *mainApp) Cancel() {
	app.canceFunc()
}

func (app *mainApp) Close() {
	app.Application.Close()
}

func (app *mainApp) NewRectFrame(rect image.Rectangle) *rectframe.RectFrame {
	return rectframe.New(app.Application, rect)
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
		e, ok := ctxchan.Get(ctx, app.Application.Chan())
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
