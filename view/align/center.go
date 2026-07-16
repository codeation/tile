package align

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// Alignment flags
const (
	_ int = iota
	Left
	Center
	Right
)

// AlignView implements horizontal alignment child view in specified width
type AlignView struct {
	view.Viewer
	align int
	width fn.Int
}

// New returns a created AlignView
func New(viewer view.Viewer) *AlignView {
	return &AlignView{
		Viewer: viewer,
		align:  Left,
	}
}

// Align sets alignment flag
func (v *AlignView) Align(align int) *AlignView {
	v.align = align
	return v
}

// Widths set horizontal width
func (v *AlignView) Width(width fn.Int) *AlignView {
	v.width = width
	return v
}

// Size returns a element drawing size
func (v *AlignView) Size() image.Point {
	elemSize := v.Viewer.Size()
	if v.width != nil {
		return image.Pt(max(v.width(), elemSize.X), elemSize.Y)
	}
	return elemSize
}

// Draw draws a element in a window width specified offset
func (v *AlignView) Draw(w *impress.Window, from image.Point) {
	elemSize := v.Viewer.Size()
	if v.width != nil {
		switch v.align {
		case Center:
			from.X += max(0, (v.width()-elemSize.X)/2)
		case Right:
			from.X += max(0, v.width()-elemSize.X)
		}
	}
	v.Viewer.Draw(w, from)
}

// Select returns active element and its rect for the click point
func (v *AlignView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	elemSize := v.Viewer.Size()
	if v.width != nil {
		switch v.align {
		case Center:
			from.X += max(0, (v.width()-elemSize.X)/2)
		case Right:
			from.X += max(0, v.width()-elemSize.X)
		}
	}
	return v.Viewer.Select(pt, from)
}
