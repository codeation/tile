package margin

import (
	"image"

	"github.com/codeation/impress"

	"github.com/codeation/tile/view"
)

// MarginView adds a margin space around any Viewer
type MarginView struct {
	view.Viewer
	margin image.Point
}

// New creates a MarginView
func New(viewer view.Viewer) *MarginView {
	return &MarginView{
		Viewer: viewer,
	}
}

// Margin sets margin space function
func (v *MarginView) Margin(margin image.Point) *MarginView {
	v.margin = margin
	return v
}

// Size returns size of a view element
func (v *MarginView) Size() image.Point {
	return v.Viewer.Size().Add(v.margin.Mul(2))
}

// Draw draws a element in a window width specified offset
func (v *MarginView) Draw(w *impress.Window, from image.Point) {
	v.Viewer.Draw(w, from.Add(v.margin))
}

// Select returns active element and its rect for the click point
func (v *MarginView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	return v.Viewer.Select(pt, from.Add(v.margin))
}
