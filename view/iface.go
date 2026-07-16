package view

import (
	"image"

	"github.com/codeation/impress"
)

// Viewer is an interface of any view element
type Viewer interface {
	// Size returns a element drawing size
	Size() image.Point
	// Draw draws any element in a window width specified offset
	Draw(w *impress.Window, from image.Point)
	// Select returns active element and its rect for the click point
	Select(pt image.Point, from image.Point) (any, image.Rectangle)
}
