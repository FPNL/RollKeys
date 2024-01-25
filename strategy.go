package rollKeys

import (
	"sync/atomic"
)

var circularIdx int64

func DefaultPickStrategy[T any](slice []*Limiter[T]) *Limiter[T] {
	atomic.AddInt64(&circularIdx, 1)

	if circularIdx >= int64(len(slice)) {
		circularIdx = 0
	}

	return slice[circularIdx]
}
