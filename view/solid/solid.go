package solid

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view/fn"
)

// Solid is a colored rectangular element
type Solid struct {
	size       image.Point
	foreground fn.Color
}

// New creates a Solid
func New(size image.Point, foreground fn.Color) *Solid {
	return &Solid{
		size:       size,
		foreground: foreground,
	}
}

// Size returns size of a view element
func (s *Solid) Size() image.Point {
	return s.size
}

// Draw draws a element in a window width specified offset
func (s *Solid) Draw(w *impress.Window, from image.Point) {
	w.Fill(image.Rectangle{Min: from, Max: from.Add(s.size)}, s.foreground())
}

// Select returns active element and its rect for the click point
func (s *Solid) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	return nil, image.Rectangle{}
}
