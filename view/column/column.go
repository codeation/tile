package column

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view"
)

// ColumnView combines elems into Viewer
type ColumnView struct {
	elems  []view.Viewer
	margin int
}

// New creates a ColumnView
func New(elems ...view.Viewer) *ColumnView {
	return &ColumnView{
		elems: elems,
	}
}

// Margin sets a margin space function
func (v *ColumnView) Margin(margin int) *ColumnView {
	v.margin = margin
	return v
}

// Size returns size of a view element
func (v *ColumnView) Size() image.Point {
	if len(v.elems) == 0 {
		return image.Point{}
	}
	output := image.Pt(0, -v.margin)
	for _, elem := range v.elems {
		elemSize := elem.Size()
		output.X = max(output.X, elemSize.X)
		output.Y += elemSize.Y + v.margin
	}
	return output
}

// Draw draws a element in a window width specified offset
func (v *ColumnView) Draw(w *impress.Window, from image.Point) {
	for _, elem := range v.elems {
		elemSize := elem.Size()
		elem.Draw(w, from)
		from.Y += elemSize.Y + v.margin
	}
}

// Select returns active element and its rect for the click point
func (v *ColumnView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	for _, elem := range v.elems {
		elemSize := elem.Size()
		if pt.In(image.Rectangle{Min: from, Max: from.Add(elemSize)}) {
			return elem.Select(pt, from)
		}
		from.Y += elemSize.Y + v.margin
	}
	return nil, image.Rectangle{}
}
