package box

import (
	"image"

	"github.com/codeation/impress"

	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/border"
	"github.com/codeation/tile/view/fn"
)

// BoxView adds a border, background color, or padding to any Viewer.
type BoxView struct {
	viewer     view.Viewer
	padding    image.Point
	foreground fn.Color
	background fn.Color
}

// New returns a BoxView for specified Viewer.
// Foreground is a border color, Background is a new background color for a Viewer.
func New(viewer view.Viewer) *BoxView {
	return &BoxView{
		viewer: viewer,
	}
}

// Padding sets padding size function.
func (v *BoxView) Padding(padding image.Point) *BoxView {
	v.padding = padding
	return v
}

// Foreground sets func for border foreground color.
func (v *BoxView) Foreground(foreground fn.Color) *BoxView {
	v.foreground = foreground
	return v
}

// Background sets func for box foreground color.
func (v *BoxView) Background(background fn.Color) *BoxView {
	v.background = background
	return v
}

// Size returns size of a view element.
func (v *BoxView) Size() image.Point {
	return v.viewer.Size().Add(v.padding.Mul(2))
}

// Draw draws a element in a window width specified offset.
func (v *BoxView) Draw(w *impress.Window, from image.Point) {
	boxRect := image.Rectangle{Min: from, Max: from.Add(v.viewer.Size().Add(v.padding.Mul(2)))}
	if v.background != nil {
		w.Fill(boxRect, v.background())
	}
	if v.foreground != nil {
		border.Border(w, boxRect, v.foreground())
	}
	v.viewer.Draw(w, from.Add(v.padding))
}

// Select returns active element and its rect for the click point.
func (v *BoxView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	return v.viewer.Select(pt, from.Add(v.padding))
}
