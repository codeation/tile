package eventlink

import (
	"context"
	"sync"

	"github.com/codeation/impress/event"

	"github.com/codeation/tile/eventlink/ctxchan"
)

// EventLink is the event channel for the child window controller
type EventLink struct {
	events     chan event.Eventer
	appFramer  AppFramer
	actor      Actor
	ctx        context.Context
	cancelFunc context.CancelFunc
	mutex      sync.RWMutex
	wg         sync.WaitGroup
}

// New creates an empty EventLink
func New(appFramer AppFramer) *EventLink {
	return &EventLink{
		events:    make(chan event.Eventer, 64),
		appFramer: appFramer,
	}
}

// Close cancels the child context and waits for the child goroutines
func (c *EventLink) Close() {
	c.Cancel()
	c.wg.Wait()
}

// Chan returns a channel for putting events to child window controller
func (c *EventLink) Chan() <-chan event.Eventer {
	return c.events
}

// Link launches the child controller
func (c *EventLink) Link(parentCtx context.Context, child Actor) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cancelFunc != nil {
		c.cancelFunc()
	}

	ctx, cancelFunc := context.WithCancel(parentCtx)
	c.actor = child
	c.ctx = ctx
	c.cancelFunc = cancelFunc

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		child.Action(ctx, AppWithLink(c.appFramer, c))
		cancelFunc()
		child.Wait()
	}()
}

// Cancel cancels the child context
func (c *EventLink) Cancel() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.cancelFunc != nil {
		c.cancelFunc()
	}
}

// Put puts the event into the child channel if context is not canceled
func (c *EventLink) Put(ctx context.Context, ev event.Eventer) {
	ctxchan.Put(ctx, c.events, ev)
}

// Actor returns child controller and child context status
func (c *EventLink) Actor() (Actor, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.actor, c.ctx != nil && c.ctx.Err() == nil
}
