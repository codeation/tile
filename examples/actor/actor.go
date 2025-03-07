package main

import (
	"context"
	"image"
	"image/color"

	_ "github.com/codeation/impress/duo"

	"github.com/codeation/impress"
	"github.com/codeation/impress/event"
	"github.com/codeation/tile/eventlink"
	"github.com/codeation/tile/eventlink/ctxchan"
)

var (
	blackColor = color.RGBA{A: 255}
	whileColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	redColor   = color.RGBA{R: 255, A: 255}
)

type SimpleActor struct{}

func (a *SimpleActor) Action(ctx context.Context, app eventlink.App) {
	font := app.Application().NewFont(15, map[string]string{"family": "Verdana"})
	defer font.Close()

	w := app.NewWindow(app.InnerRect(), whileColor)
	defer w.Drop()

	for {
		if len(app.Chan()) == 0 {
			drawMessage(app, w, font, "Hello, world!")
		}

		e, ok := ctxchan.Get(ctx, app.Chan())
		if !ok {
			return
		}

		switch ev := e.(type) {
		case event.General:
			switch ev {
			case event.DestroyEvent:
				app.Cancel()
				return
			}

		case event.Keyboard:
			switch ev {
			case event.KeyEscape, event.KeyExit:
				app.Cancel()
				return
			}

		case event.Configure:
			w.Size(app.InnerRect())
		}
	}
}

func (a *SimpleActor) Wait() {}

func drawMessage(app eventlink.App, w *impress.Window, font *impress.Font, message string) {
	w.Clear()
	textSize := font.Size(message)
	offset := app.InnerRect().Size().Sub(textSize).Div(2)
	w.Text(message, font, offset, blackColor)
	w.Line(offset.Add(image.Pt(0, textSize.Y)), offset.Add(textSize), redColor)
	w.Show()
	app.Application().Sync()
}

func main() {
	app := eventlink.MainApp(impress.NewApplication(image.Rect(0, 0, 480, 240), "Panels"))
	defer app.Close()

	app.Run(context.Background(), new(SimpleActor))
}
