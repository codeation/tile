package syncvar

import (
	"sync"
)

// Var contains a variable of any type protected by a mutex
type Var[T any] struct {
	value T
	mutex sync.RWMutex
}

// New creates a new Var with specified value
func New[T any](value T) *Var[T] {
	return &Var[T]{
		value: value,
	}
}

// Zero creates a new Var with a zero value
func Zero[T any]() *Var[T] {
	return &Var[T]{}
}

// Set sets a value
func (v *Var[T]) Set(value T) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.value = value
}

// Get returns a value
func (v *Var[T]) Get() T {
	v.mutex.RLock()
	defer v.mutex.RUnlock()
	return v.value
}

// Swap sets a new value and returns the previous value
func (v *Var[T]) Swap(value T) T {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	output := v.value
	v.value = value
	return output
}
