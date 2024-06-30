package foreword

import (
	"image"

	"github.com/codeation/impress"
	"github.com/codeation/tile/view/fn"
)

// ForewordView draws a first lines of large text
type ForewordView struct {
	text       fn.String
	font       *impress.Font
	lineHeight int
	foreground fn.Color
	lineCount  int
}

// New creates a ForewordView
func New(text fn.String, font *impress.Font, lineHeight int, foreground fn.Color, lineCount int) *ForewordView {
	return &ForewordView{
		text:       text,
		font:       font,
		lineHeight: lineHeight,
		foreground: foreground,
		lineCount:  lineCount,
	}
}

func (v *ForewordView) fontDelta() int {
	return v.lineHeight - v.font.Height
}

// Size returns size of a view element. Width of size parameter is used to split a text into sublines
func (v *ForewordView) Size(size image.Point) image.Point {
	lineCount := 0
	for _, text := range Split(v.text()) {
		lineCount += len(v.font.Split(text, size.X, 0))
		if lineCount > v.lineCount {
			lineCount = v.lineCount
			break
		}
	}
	lineCount = max(lineCount, 1)
	return image.Pt(size.X, v.lineHeight*lineCount-v.fontDelta())
}

// Draw draws a view element. Width of rect parameter is used to split a text into sublines
func (v *ForewordView) Draw(w *impress.Window, rect image.Rectangle) {
	from := rect.Min
	lineCount := 0
	for _, text := range Split(v.text()) {
		lines := v.font.Split(text, rect.Dx(), 0)
		for _, line := range lines {
			w.Text(line, v.font, from, v.foreground())
			from.Y += v.lineHeight
			lineCount++
			if lineCount >= v.lineCount {
				return
			}
		}
	}
}
