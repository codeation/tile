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
	blackColor color.Color = color.RGBA{A: 255}
	whileColor color.Color = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	redColor   color.Color = color.RGBA{R: 255, A: 255}
)

type Example struct {
	font *impress.Font
}

func (e *Example) Action(ctx context.Context, app eventlink.App) {
	w := app.NewWindow(app.InnerRect(), whileColor)
	defer w.Drop()

	for {
		if len(app.Chan()) == 0 {
			w.Clear()
			w.Text("Hello, world!", e.font, image.Pt(200, 100), blackColor)
			w.Line(image.Pt(200, 120), image.Pt(300, 120), redColor)
			w.Show()
			app.Application().Sync()
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

		case event.Configure:
			w.Size(app.InnerRect())

		}
	}
}

func (e *Example) Wait() {}

func main() {
	ctx := context.Background()

	rect := image.Rect(0, 0, 480, 240)
	app := eventlink.MainApp(impress.NewApplication(rect, "Panels"))
	defer app.Close()

	exampleActor := &Example{font: app.Application().NewFont(15, map[string]string{"family": "Verdana"})}
	defer exampleActor.font.Close()
	app.Run(ctx, exampleActor)
}
