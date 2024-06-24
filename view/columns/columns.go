package columns

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// ColumnsView combines elems into Viewer
type ColumnsView struct {
	elems   []view.Viewer
	marginX fn.Int
}

// New creates a ColumnView
func New(marginX fn.Int, elems ...view.Viewer) *ColumnsView {
	return &ColumnsView{
		elems:   elems,
		marginX: marginX,
	}
}

// Size returns size of a view element
func (v *ColumnsView) Size(size image.Point) image.Point {
	if len(v.elems) == 0 {
		return image.Point{}
	}
	marginX := v.marginX()
	output := image.Pt(-marginX, 0)
	for _, elem := range v.elems {
		elemSize := elem.Size(size)
		output.X += elemSize.X + marginX
		output.Y = max(output.Y, elemSize.Y)
	}
	return output
}

// Draw draws a view element
func (v *ColumnsView) Draw(w *impress.Window, rect image.Rectangle) {
	marginX := v.marginX()
	for _, elem := range v.elems {
		elemSize := elem.Size(rect.Size())
		elem.Draw(w, rect)
		rect.Min.X += elemSize.X + marginX
	}
}
