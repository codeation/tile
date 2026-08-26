// Package implements array of control elements
package sequence

import (
	"context"
	"slices"

	"github.com/codeation/impress/event"
	"github.com/codeation/tile/control"
	"github.com/codeation/tile/eventlink"
)

// Sequence contains a slice of control elemens and selected element.
type Sequence struct {
	elems            []any
	selected         any
	prevKey, nextKey event.Eventer
}

// New creates a empty vertical oriented sequence.
func New() *Sequence {
	return &Sequence{
		prevKey: event.KeyUp,
		nextKey: event.KeyDown,
	}
}

// Add adds a one element to sequence.
func (c *Sequence) Add(elem any) *Sequence {
	if c.selected == nil {
		c.selected = elem
	}
	c.elems = append(c.elems, elem)
	return c
}

// Vertical changes sequence to vertically arranged.
func (c *Sequence) Vertical() *Sequence {
	c.prevKey = event.KeyUp
	c.nextKey = event.KeyDown
	return c
}

// Vertical changes sequence to horizontally arranged.
func (c *Sequence) Horizontal() *Sequence {
	c.prevKey = event.KeyLeft
	c.nextKey = event.KeyRight
	return c
}

// Select changes a selected element.
func (c *Sequence) Select(elem any) {
	c.selected = elem
}

// Selected returns a selected element.
func (c *Sequence) Selected() any {
	return c.selected
}

// Control implements a reaction to one event.
func (c *Sequence) Control(ctx context.Context, app eventlink.App, e event.Eventer, prior control.DoFunc) {
	switch e {
	case c.prevKey:
		if elem := c.Prev(c.Selected()); elem != nil {
			c.selected = elem
			return
		}
		prior(ctx, app, e)
	case c.nextKey:
		if elem := c.Next(c.Selected()); elem != nil {
			c.selected = elem
			return
		}
		prior(ctx, app, e)
	default:
		if elemControl, ok := c.selected.(control.Control); ok {
			elemControl.Control(ctx, app, e, prior)
			return
		}
		prior(ctx, app, e)
	}
}

// Prev returns a previous element.
func (c *Sequence) Prev(elem any) any {
	index := slices.Index(c.elems, elem)
	if index <= 0 {
		return nil
	}
	return c.elems[index-1]
}

// Next returns a next element.
func (c *Sequence) Next(elem any) any {
	index := slices.Index(c.elems, elem)
	if index < 0 || index >= len(c.elems)-1 {
		return nil
	}
	return c.elems[index+1]
}

// Front returns a first element.
func (c *Sequence) Front() any {
	if len(c.elems) != 0 {
		return c.elems[0]
	}
	return nil
}

// Back returns a latest element.
func (c *Sequence) Back() any {
	if len(c.elems) != 0 {
		return c.elems[len(c.elems)-1]
	}
	return nil
}
