package box

import (
	"image"

	"github.com/codeation/impress"

	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/border"
	"github.com/codeation/tile/view/fn"
)

// BoxView adds a border and background color to any Viewer
type BoxView struct {
	view.Viewer
	foreground fn.Color
	background fn.Color
}

// New returns a BoxView for specified Viewer. Foreground is a border color, Background is a new background color for a Viewer
func New(viewer view.Viewer, foreground fn.Color, background fn.Color) *BoxView {
	return &BoxView{
		Viewer:     viewer,
		foreground: foreground,
		background: background,
	}
}

// Draw draws a element in a window width specified offset
func (v *BoxView) Draw(w *impress.Window, from image.Point) {
	boxRect := image.Rectangle{Min: from, Max: from.Add(v.Size())}
	w.Fill(boxRect, v.background())
	border.Border(w, boxRect, v.foreground())
	v.Viewer.Draw(w, from)
}

// Select returns active element and its rect for the click point
func (v *BoxView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	return v.Viewer.Select(pt, from)
}
