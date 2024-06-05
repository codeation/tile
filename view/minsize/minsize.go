package minsize

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// MinSizeView expands a size of any Viewer to specified width
type MinSizeView struct {
	view.Viewer
	size fn.Point
}

// New returns a MinSizeView
func New(view view.Viewer, size fn.Point) *MinSizeView {
	return &MinSizeView{
		Viewer: view,
		size:   size,
	}
}

// Size returns size of a view element
func (v *MinSizeView) Size(size image.Point) image.Point {
	minSize := v.size()
	innerSize := v.Viewer.Size(minSize)
	return image.Pt(max(minSize.X, innerSize.X), max(minSize.Y, innerSize.Y))
}

// Draw draws a view element
func (v *MinSizeView) Draw(w *impress.Window, rect image.Rectangle) {
	minSize := v.size()
	v.Viewer.Draw(w, image.Rectangle{Min: rect.Min, Max: rect.Min.Add(minSize)})
}
