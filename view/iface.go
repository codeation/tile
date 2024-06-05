package view

import (
	"image"

	"github.com/codeation/impress"
)

// Viewer is an interface of any view element
type Viewer interface {
	Size(size image.Point) image.Point
	Draw(w *impress.Window, rect image.Rectangle)
}
