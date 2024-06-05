package fn

import (
	"image"
	"image/color"
)

type Int func() int
type String func() string
type Bool func() bool

type Color func() color.Color
type Point func() image.Point
type Rect func() image.Rectangle
