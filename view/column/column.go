package column

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// ColumnView combines elems into Viewer
type ColumnView struct {
	elems   []view.Viewer
	marginY fn.Int
}

// New creates a ColumnView
func New(marginY fn.Int, elems ...view.Viewer) *ColumnView {
	return &ColumnView{
		elems:   elems,
		marginY: marginY,
	}
}

// Size returns size of a view element
func (v *ColumnView) Size(size image.Point) image.Point {
	if len(v.elems) == 0 {
		return image.Point{}
	}
	marginY := v.marginY()
	output := image.Pt(0, -marginY)
	for _, elem := range v.elems {
		elemSize := elem.Size(size)
		output.X = max(output.X, elemSize.X)
		output.Y += elemSize.Y + marginY
	}
	return output
}

// Draw draws a view element
func (v *ColumnView) Draw(w *impress.Window, rect image.Rectangle) {
	marginY := v.marginY()
	for _, elem := range v.elems {
		elemSize := elem.Size(rect.Size())
		elem.Draw(w, rect)
		rect.Min.Y += elemSize.Y + marginY
	}
}
