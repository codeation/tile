package rectframe

import (
	"image"

	"github.com/codeation/impress"

	"github.com/codeation/tile/eventlink/syncvar"
)

// RectFrame contains a frame and its rectangle
type RectFrame struct {
	*impress.Frame
	rect *syncvar.Var[image.Rectangle]
}

// New creates a new frame
func New(
	framer interface {
		NewFrame(rect image.Rectangle) *impress.Frame
	},
	rect image.Rectangle,
) *RectFrame {
	return &RectFrame{
		Frame: framer.NewFrame(rect),
		rect:  syncvar.New(rect),
	}
}

// NewRectFrame creates a child frame
func (f *RectFrame) NewRectFrame(rect image.Rectangle) *RectFrame {
	return New(f.Frame, rect)
}

// Size changes the size and position of the frame
func (f *RectFrame) Size(rect image.Rectangle) {
	f.rect.Set(rect)
	f.Frame.Size(rect)
}

// Rect returns the frame coordinates
func (f *RectFrame) Rect() image.Rectangle {
	return f.rect.Get()
}

// InnerRect returns the coordinates of the inner rectangle
func (f *RectFrame) InnerRect() image.Rectangle {
	return image.Rectangle{Max: f.rect.Get().Size()}
}
