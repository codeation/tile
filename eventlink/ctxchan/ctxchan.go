package ctxchan

import (
	"context"
)

// Put puts a message into a channel unless the context is canceled
func Put[T any](ctx context.Context, c chan<- T, message T) {
	select {
	case <-ctx.Done():
		return
	case c <- message:
		return
	}
}

// Get waits for a message from the channel or returns false if the context is canceled
func Get[T any](ctx context.Context, c <-chan T) (message T, ok bool) {
	select {
	case <-ctx.Done():
		return
	case message, ok = <-c:
		return
	}
}
