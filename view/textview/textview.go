package textview

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fieldview"
	"github.com/codeation/tile/view/fn"
)

// TextView is a text element. TextView is used to draw the field's value inside any rectangle
type TextView struct {
	field      fieldview.FocusFielder
	font       *impress.Font
	lineHeight int
	foreground fn.Color
	cursor     view.Viewer
}

// New creates a TextView
func New(fielder fieldview.FocusFielder, font *impress.Font, lineHeight int, foreground fn.Color, cursor view.Viewer) *TextView {
	return &TextView{
		field:      fielder,
		font:       font,
		lineHeight: lineHeight,
		foreground: foreground,
		cursor:     cursor,
	}
}

func (v *TextView) fontDelta() int {
	return v.lineHeight - v.font.Height
}

// Size returns size of a view element. Width of size parameter is used to split a text into sublines
func (v *TextView) Size(size image.Point) image.Point {
	lineCount := len(v.font.Split(v.field.String(), size.X, 0))
	if lineCount == 0 {
		lineCount = 1
	}

	return image.Pt(size.X, v.lineHeight*lineCount-v.fontDelta())
}

// Draw draws a view element. Width of rect parameter is used to split a text into sublines
func (v *TextView) Draw(w *impress.Window, rect image.Rectangle) {
	lines := v.font.Split(v.field.String(), rect.Dx(), 0)
	from := rect.Min
	for _, line := range lines {
		w.Text(line, v.font, from, v.foreground())
		from = from.Add(image.Pt(0, v.lineHeight))
	}

	if v.field.Focused() {
		cursorIndex := v.field.Cursor()
		cursorPoint := rect.Min
		for i, line := range lines {
			if cursorIndex <= len(line) {
				cursorPoint = cursorPoint.Add(image.Pt(v.font.Size(line[:cursorIndex]).X, i*v.lineHeight))
				break
			}
			cursorIndex -= len(line)
		}
		v.cursor.Draw(w, image.Rectangle{Min: cursorPoint, Max: cursorPoint.Add(image.Pt(0, v.font.LineHeight))})
	}
}
