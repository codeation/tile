package border

import (
	"image"
	"image/color"

	"github.com/codeation/impress"
)

// ComplexBorder draws a rectangle border in specied window.
// Rect is a border coordinates, lightColor and darkColor are border segment colors, borderWidth is a border width
func ComplexBorder(w *impress.Window, rect image.Rectangle, lightColor color.Color, darkColor color.Color, borderWidth int) {
	w.Fill(image.Rect(rect.Min.X, rect.Min.Y, rect.Max.X-borderWidth, rect.Min.Y+borderWidth), lightColor)
	w.Fill(image.Rect(rect.Min.X, rect.Min.Y+borderWidth, rect.Min.X+borderWidth, rect.Max.Y), lightColor)
	w.Fill(image.Rect(rect.Max.X-borderWidth, rect.Min.Y, rect.Max.X, rect.Max.Y-borderWidth), darkColor)
	w.Fill(image.Rect(rect.Min.X+borderWidth, rect.Max.Y-borderWidth, rect.Max.X, rect.Max.Y), darkColor)
}

// Border draws a rectangle border in specied window. Rect is a border coordinates, foreground is a border line color
func Border(w *impress.Window, rect image.Rectangle, foreground color.Color) {
	ComplexBorder(w, rect, foreground, foreground, 1)
}

// InnerRect returns coordinates of an inner space rectangle inside a border
func InnerRect(rect image.Rectangle) image.Rectangle {
	return image.Rect(rect.Min.X+1, rect.Min.Y+1, rect.Max.X-1, rect.Max.Y-1)
}
