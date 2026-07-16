package rectframe

import (
	"image"
	"sync/atomic"

	"github.com/codeation/impress"
)

// RectFrame contains a frame and its rectangle
type RectFrame struct {
	*impress.Frame
	rect atomic.Value // image.Rectangle
}

// New creates a new frame
func New(
	framer interface {
		NewFrame(rect image.Rectangle) *impress.Frame
	},
	rect image.Rectangle,
) *RectFrame {
	f := &RectFrame{
		Frame: framer.NewFrame(rect),
	}
	f.rect.Store(rect)
	return f
}

// NewRectFrame creates a child frame
func (f *RectFrame) NewRectFrame(rect image.Rectangle) *RectFrame {
	return New(f.Frame, rect)
}

// Size changes the size and position of the frame
func (f *RectFrame) Size(rect image.Rectangle) {
	f.rect.Store(rect)
	f.Frame.Size(rect)
}

// Rect returns the frame coordinates
func (f *RectFrame) Rect() image.Rectangle {
	if v, ok := f.rect.Load().(image.Rectangle); ok {
		return v
	}
	return image.Rectangle{}
}

// InnerRect returns the coordinates of the inner rectangle
func (f *RectFrame) InnerRect() image.Rectangle {
	if v, ok := f.rect.Load().(image.Rectangle); ok {
		return image.Rectangle{Max: v.Size()}
	}
	return image.Rectangle{}
}
