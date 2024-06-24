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
func (*EmptyView) Size(size image.Point) image.Point { return image.Point{} }

// Draw draws a view element
func (*EmptyView) Draw(w *impress.Window, rect image.Rectangle) {}
