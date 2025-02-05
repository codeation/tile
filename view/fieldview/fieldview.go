package fieldview

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/elem/nl"
	"github.com/codeation/tile/view"
	"github.com/codeation/tile/view/fn"
)

// FieldView is a text element. FieldView is used to draw the field's value inside any rectangle
type FieldView struct {
	field      FocusFielder
	font       *impress.Font
	foreground fn.Color
	lineHeight int
	cursor     view.Viewer
	maxRows    int
}

// New creates a FieldView
func New(fielder FocusFielder, font *impress.Font, foreground fn.Color) *FieldView {
	return &FieldView{
		field:      fielder,
		font:       font,
		foreground: foreground,
		lineHeight: font.LineHeight,
	}
}

// LineHeight sets a custom line height
func (v *FieldView) LineHeight(lineHeight int) *FieldView {
	v.lineHeight = lineHeight
	return v
}

// WithCursor adds a cursor Viewer to draw in cursor position
func (v *FieldView) WithCursor(cursor view.Viewer) *FieldView {
	v.cursor = cursor
	return v
}

// MaxRows sets a maximum row height
func (v *FieldView) MaxRows(maxRows int) *FieldView {
	v.maxRows = maxRows
	return v
}

func (v *FieldView) splitText(rect image.Rectangle) ([]string, int, int) {
	output := []string{}
	cursorIndex := v.field.Cursor()
	row := 0
	col := 0
	for _, paragraph := range v.field.Strings() {
		lines := v.font.Split(paragraph, rect.Dx(), 0)
		switch {
		case cursorIndex > len(paragraph)+len(nl.DefaultNewLine.String()):
			cursorIndex -= len(paragraph) + len(nl.DefaultNewLine.String())
		case cursorIndex == len(paragraph)+len(nl.DefaultNewLine.String()):
			row = len(output) + len(lines)
			cursorIndex = 0
		case cursorIndex >= 0:
			for i, line := range lines {
				if cursorIndex <= len(line) {
					row = len(output) + i
					col = cursorIndex
					cursorIndex = -1
					break
				}
				cursorIndex -= len(line)
			}
		}
		output = append(output, lines...)
	}
	if v.maxRows != 0 {
		halfRow := min(row+(v.maxRows+1)/2, len(output))
		minRow := max(halfRow-v.maxRows, 0)
		maxRow := min(minRow+v.maxRows, len(output))
		output = output[minRow:maxRow]
		row -= minRow
	}
	return output, row, col
}

// Size returns size of a view element. Width of size parameter is used to split a text into sublines
func (v *FieldView) Size(size image.Point) image.Point {
	lineCount := 0
	for _, s := range v.field.Strings() {
		lineCount += len(v.font.Split(s, size.X, 0))
		if v.maxRows != 0 && lineCount > v.maxRows {
			break
		}
	}

	lineCount = max(lineCount, 1)
	if v.maxRows != 0 {
		lineCount = min(lineCount, v.maxRows)
	}

	return image.Pt(size.X, v.lineHeight*lineCount-(v.lineHeight-v.font.Height))
}

// Draw draws a view element. Width of rect parameter is used to split a text into sublines
func (v *FieldView) Draw(w *impress.Window, rect image.Rectangle) {
	lines, row, col := v.splitText(rect)
	from := rect.Min
	var cursorPoint image.Point
	for i, line := range lines {
		w.Text(line, v.font, from, v.foreground())
		if i == row {
			cursorPoint = from
		}
		from = from.Add(image.Pt(0, v.lineHeight))
	}

	if v.cursor != nil && v.field.Focused() {
		cursorPoint = rect.Min.Add(image.Pt(v.font.Size(lines[row][:col]).X, row*v.lineHeight))
		v.cursor.Draw(w, image.Rectangle{Min: cursorPoint, Max: cursorPoint.Add(image.Pt(0, v.font.LineHeight))})
	}
}
