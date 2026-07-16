package row

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// RowView combines columns into Viewer
type RowView struct {
	elems  []view.Viewer
	margin int
	width  fn.Int
}

// New creates a RowView
func New(elems ...view.Viewer) *RowView {
	return &RowView{
		elems: elems,
	}
}

// Width sets a fixed width for element
func (v *RowView) Width(width fn.Int) *RowView {
	v.width = width
	return v
}

// MarginX sets a margin space function
func (v *RowView) Margin(margin int) *RowView {
	v.margin = margin
	return v
}

// Size returns size of a view element
func (v *RowView) Size() image.Point {
	if len(v.elems) == 0 {
		return image.Point{}
	}
	output := image.Pt(-v.margin, 0)
	for _, elem := range v.elems {
		elemSize := elem.Size()
		if v.width != nil && v.width() < output.X+elemSize.X+v.margin {
			break
		}
		output.X += elemSize.X + v.margin
		output.Y = max(output.Y, elemSize.Y)
	}
	return output
}

// Draw draws a element in a window width specified offset
func (v *RowView) Draw(w *impress.Window, from image.Point) {
	offset := from
	for _, elem := range v.elems {
		elemSize := elem.Size()
		if v.width != nil && v.width() < offset.X+elemSize.X-from.X {
			break
		}
		elem.Draw(w, offset)
		offset.X += elemSize.X + v.margin
	}
}

// Select returns active element and its rect for the click point
func (v *RowView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	offset := from
	for _, elem := range v.elems {
		elemSize := elem.Size()
		if v.width != nil && v.width() < offset.X+elemSize.X-from.X {
			break
		}
		if pt.In(image.Rectangle{Min: offset, Max: offset.Add(elemSize)}) {
			return elem.Select(pt, offset)
		}
		offset.X += elemSize.X + v.margin
	}
	return nil, image.Rectangle{}
}
