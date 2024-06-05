package margin

import (
	"image"

	"github.com/codeation/impress"

	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// MarginView adds a margin space around any Viewer
type MarginView struct {
	view.Viewer
	margin fn.Point
}

// New creates a MarginView
func New(viewer view.Viewer, margin fn.Point) *MarginView {
	return &MarginView{
		Viewer: viewer,
		margin: margin,
	}
}

// Size returns size of a view element
func (v *MarginView) Size(size image.Point) image.Point {
	marginSize := v.margin()
	innerSize := v.Viewer.Size(size)
	return innerSize.Add(marginSize.Mul(2))
}

// Draw draws a view element
func (v *MarginView) Draw(w *impress.Window, rect image.Rectangle) {
	marginSize := v.margin()
	innerRect := image.Rectangle{Min: rect.Min.Add(marginSize), Max: rect.Max.Sub(marginSize)}
	v.Viewer.Draw(w, innerRect)
}
