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

// Draw draws a view element
func (v *BoxView) Draw(w *impress.Window, rect image.Rectangle) {
	size := v.Viewer.Size(rect.Size())
	boxRect := image.Rectangle{Min: rect.Min, Max: rect.Min.Add(size)}
	w.Fill(boxRect, v.background())
	border.Border(w, boxRect, v.foreground())
	v.Viewer.Draw(w, rect)
}
