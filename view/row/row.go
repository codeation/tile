package row

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// RowView combines columns into Viewer
type RowView struct {
	elems   []view.Viewer
	marginX fn.Int
}

// New creates a RowView
func New(elems ...view.Viewer) *RowView {
	return &RowView{
		elems:   elems,
		marginX: fn.Const(0),
	}
}

// MarginX sets a margin space function
func (v *RowView) MarginX(marginX fn.Int) *RowView {
	v.marginX = marginX
	return v
}

// Size returns size of a view element
func (v *RowView) Size(size image.Point) image.Point {
	if len(v.elems) == 0 {
		return image.Pt(size.X, 0)
	}
	marginX := v.marginX()
	x, y := -marginX, 0
	for _, elem := range v.elems {
		elemSize := elem.Size(size)
		if x+elemSize.X+marginX > size.X {
			break
		}
		x += elemSize.X + marginX
		y = max(y, elemSize.Y)
	}
	return image.Pt(max(x, size.X), y)
}

// Draw draws a view element
func (v *RowView) Draw(w *impress.Window, rect image.Rectangle) {
	marginX := v.marginX()
	for _, elem := range v.elems {
		elemSize := elem.Size(rect.Size())
		if rect.Min.X+elemSize.X > rect.Max.X {
			break
		}
		elem.Draw(w, rect)
		rect.Min.X += elemSize.X + marginX
	}
}
