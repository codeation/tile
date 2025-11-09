package eventlink

import (
	"context"
	"sync"

	"github.com/codeation/impress/event"

	"github.com/codeation/tile/eventlink/ctxchan"
)

// EventLink represents a channel for managing events between the parent and child controllers.
type EventLink struct {
	events     chan event.Eventer
	appFramer  AppFramer
	actor      Actor
	ctx        context.Context
	cancelFunc context.CancelFunc
	mutex      sync.RWMutex
	wg         sync.WaitGroup
}

// New creates and returns a new EventLink.
func New() *EventLink { return new(EventLink) }

// Close cancels the child context and waits for all child goroutines to complete.
func (c *EventLink) Close() {
	c.Cancel()
	c.wg.Wait()
}

// Chan returns a channel for receiving events intended for the child window controller.
func (c *EventLink) Chan() <-chan event.Eventer {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.events
}

// Link launches the child window controller. Previous child controller context will be cancelled, if one exists.
func (c *EventLink) Link(parentCtx context.Context, appFramer AppFramer, child Actor) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.cancelFunc != nil {
		c.cancelFunc()
	}

	c.events = make(chan event.Eventer, 64)
	c.appFramer = appFramer

	ctx, cancelFunc := context.WithCancel(parentCtx)
	c.actor = child
	c.ctx = ctx
	c.cancelFunc = cancelFunc

	c.wg.Go(func() {
		child.Action(ctx, AppWithLink(c.appFramer, c))
		cancelFunc()
		child.Wait()
	})
}

// Cancel cancels the current child context, if one exists.
func (c *EventLink) Cancel() {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.cancelFunc != nil {
		c.cancelFunc()
	}
}

// Put sends an event to the child channel if the context has not been canceled.
func (c *EventLink) Put(ctx context.Context, ev event.Eventer) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	ctxchan.Put(ctx, c.events, ev)
}

// Actor returns the current child controller and a boolean indicating if the child context is still active.
func (c *EventLink) Actor() (Actor, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.actor, c.ctx != nil && c.ctx.Err() == nil
}
