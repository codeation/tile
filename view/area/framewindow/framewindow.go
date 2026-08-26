package framewindow

import (
	"image/color"

	"github.com/codeation/impress"
	"github.com/codeation/tile/eventlink"
)

// FrameWindow is a full-size window container that implements area.Area interface.
type FrameWindow struct {
	w   *impress.Window
	app *impress.Application
}

// New creates a full-size window container.
func New(app eventlink.AppFramer, background color.Color) *FrameWindow {
	return &FrameWindow{
		w:   app.NewWindow(app.InnerRect(), background),
		app: app.Application(),
	}
}

// Window returns a container window.
func (fw *FrameWindow) Window() *impress.Window { return fw.w }

// Configure reacts to event.Configure.
func (fw *FrameWindow) Configure(app eventlink.AppFramer) { fw.w.Size(app.InnerRect()) }

// Drop drops a container.
func (fw *FrameWindow) Drop() { fw.w.Drop(); fw.app.Sync() }
