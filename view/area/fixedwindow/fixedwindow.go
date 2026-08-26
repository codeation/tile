package fixedwindow

import (
	"image"
	"image/color"

	"github.com/codeation/impress"
	"github.com/codeation/tile/eventlink"
)

// FixedWindow is a fixed-size window that implements area.Area interface.
type FixedWindow struct {
	w   *impress.Window
	app *impress.Application
}

// New creates a fixed-size window container.
func New(app eventlink.AppFramer, rect image.Rectangle, background color.Color) *FixedWindow {
	return &FixedWindow{
		w:   app.NewWindow(rect, background),
		app: app.Application(),
	}
}

// Window returns a container window.
func (fw *FixedWindow) Window() *impress.Window { return fw.w }

// Configure reacts to event.Configure.
func (fw *FixedWindow) Configure(app eventlink.AppFramer) {}

// Drop drops a container.
func (fw *FixedWindow) Drop() { fw.w.Drop(); fw.app.Sync() }
