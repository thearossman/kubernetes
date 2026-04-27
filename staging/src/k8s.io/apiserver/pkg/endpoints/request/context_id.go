package request

import (
	"context"
	"sync/atomic"
)

// To attach value to context, need to give it a type. 
// See the ServerShutdown key type example.
// Needs to be used/imported anywhere that needs to use it.
type contextIDKeyType int

const contextIDKey contextIDKeyType = iota

var requestIDCounter uint64

func ContextIDFrom(ctx context.Context) (uint64, bool) {
	id, ok := ctx.Value(contextIDKey).(uint64)
	return id, ok
}

func WithContextID(parent context.Context) (context.Context, uint64) {
	if id, ok := ContextIDFrom(parent); ok {
		return parent, id
	}
	id := atomic.AddUint64(&requestIDCounter, 1)
	return context.WithValue(parent, contextIDKey, id), id
}

