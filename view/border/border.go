package border

import (
	"image"
	"image/color"

	"github.com/codeation/impress"
)

// Border draws a rectangle border in specied window. Rect is a border coordinates, foreground is a border line color
func Border(w *impress.Window, rect image.Rectangle, foreground color.Color) {
	w.Line(image.Pt(rect.Min.X+1, rect.Min.Y), image.Pt(rect.Max.X, rect.Min.Y), foreground)
	w.Line(image.Pt(rect.Max.X-1, rect.Min.Y+1), image.Pt(rect.Max.X-1, rect.Max.Y), foreground)
	w.Line(image.Pt(rect.Min.X, rect.Max.Y-1), image.Pt(rect.Max.X-1, rect.Max.Y-1), foreground)
	w.Line(image.Pt(rect.Min.X, rect.Min.Y), image.Pt(rect.Min.X, rect.Max.Y-1), foreground)
}

// InnerRect returns coordinates of an inner space rectangle inside a border
func InnerRect(rect image.Rectangle) image.Rectangle {
	return image.Rect(rect.Min.X+1, rect.Min.Y+1, rect.Max.X-1, rect.Max.Y-1)
}
