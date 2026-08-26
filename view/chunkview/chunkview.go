package chunkview

import (
	"image"

	"github.com/codeation/tile/view"
)

// ChunkView is a container that returns specified model as select result.
type ChunkView struct {
	view.Viewer
	model any
}

// New creates ChunkView.
func New(child view.Viewer) *ChunkView {
	return &ChunkView{
		Viewer: child,
	}
}

// Model specifies model.
func (v *ChunkView) Model(m any) *ChunkView {
	v.model = m
	return v
}

// Select returns a model when a click point inside view rectangle.
func (v *ChunkView) Select(pt image.Point, from image.Point) (any, image.Rectangle) {
	rect := image.Rectangle{Max: v.Viewer.Size()}.Add(from)
	if pt.In(rect) {
		return v.model, rect
	}
	return nil, image.Rectangle{}
}
