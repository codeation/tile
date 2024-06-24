package fieldview

import (
	"image"

	"github.com/codeation/impress"

	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// Fielder is an interface of a field model
type Fielder interface {
	String() string
	Cursor() int
}

// FocusFielder is an interface of a field model with a focused flag
type FocusFielder interface {
	Fielder
	Focused() bool
}

// FieldView is a Viewer for a field. Cursor is an element for drawing in selected field
type FieldView struct {
	field      FocusFielder
	font       *impress.Font
	foreground fn.Color
	cursor     view.Viewer
}

// New creates a FieldView
func New(fielder FocusFielder, font *impress.Font, foreground fn.Color, cursor view.Viewer) *FieldView {
	return &FieldView{
		field:      fielder,
		font:       font,
		foreground: foreground,
		cursor:     cursor,
	}
}

// Size returns size of a view element
func (v *FieldView) Size(size image.Point) image.Point {
	return v.font.Size(v.field.String()).Add(image.Pt(v.cursor.Size(image.Point{}).X, 0))
}

// Draw draws a view element
func (v *FieldView) Draw(w *impress.Window, rect image.Rectangle) {
	text := v.field.String()
	w.Text(text, v.font, rect.Min, v.foreground())

	if v.field.Focused() {
		cursorPoint := rect.Min.Add(image.Pt(v.font.Size(text[:v.field.Cursor()]).X, 0))
		v.cursor.Draw(w, image.Rectangle{Min: cursorPoint, Max: cursorPoint.Add(image.Pt(0, v.font.LineHeight))})
	}
}
