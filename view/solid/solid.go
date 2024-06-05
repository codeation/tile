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
func (s *Solid) Size(size image.Point) image.Point {
	return s.size
}

// Draw draws a view element
func (s *Solid) Draw(w *impress.Window, rect image.Rectangle) {
	w.Fill(image.Rectangle{Min: rect.Min, Max: rect.Min.Add(s.size)}, s.foreground())
}
