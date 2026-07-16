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
	width      fn.Int
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

// Width sets a func to return width of element
func (v *FieldView) Width(width fn.Int) *FieldView {
	v.width = width
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

func (v *FieldView) splitText() ([]string, int, int) {
	output := []string{}
	cursorIndex := v.field.Cursor()
	row := 0
	col := 0
	for _, paragraph := range v.field.Strings() {
		var lines []string
		if v.width != nil {
			lines = v.font.Split(paragraph, v.width(), 0)
		} else {
			lines = append(lines, paragraph)
		}
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

// Size returns a element drawing size
func (v *FieldView) Size() image.Point {
	maxWidth := 0
	if v.width != nil {
		maxWidth = v.width()
	}
	lineCount := 0
	for _, s := range v.field.Strings() {
		if v.width != nil {
			lineCount += len(v.font.Split(s, v.width(), 0))
		} else {
			lineCount++
			maxWidth = max(maxWidth, v.font.Size(s).X)
		}
		if v.maxRows != 0 && lineCount > v.maxRows {
			break
		}
	}

	lineCount = max(lineCount, 1)
	if v.maxRows != 0 {
		lineCount = min(lineCount, v.maxRows)
	}

	return image.Pt(maxWidth, v.lineHeight*lineCount-(v.lineHeight-v.font.Height))
}

// Draw draws a element in a window width specified offset
func (v *FieldView) Draw(w *impress.Window, from image.Point) {
	lines, row, col := v.splitText()
	pt := from
	for _, line := range lines {
		w.Text(line, v.font, pt, v.foreground())
		pt.Y += v.lineHeight
	}

	if v.cursor != nil && v.field.Focused() {
		cursorPoint := from.Add(image.Pt(v.font.Size(lines[row][:col]).X, row*v.lineHeight))
		v.cursor.Draw(w, cursorPoint)
	}
}

// Select returns active element and its rect for the click point
func (v *FieldView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	elemRect := image.Rectangle{Min: from, Max: from.Add(v.Size())}
	if !pt.In(elemRect) {
		return nil, image.Rectangle{}
	}
	return v.field, elemRect
}
