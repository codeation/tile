package empty

import (
	"image"

	"github.com/codeation/impress"
)

// EmptyView contains no draw elements
type EmptyView struct{}

// New creates a FieldView
func New() *EmptyView { return &EmptyView{} }

// Size returns size of a view element
func (*EmptyView) Size() image.Point { return image.Point{} }

// Draw draws a element in a window width specified offset
func (*EmptyView) Draw(w *impress.Window, from image.Point) {}

// Select returns active element and its rect for the click point
func (EmptyView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	return nil, image.Rectangle{}
}
