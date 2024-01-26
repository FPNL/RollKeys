package rollKeys

import (
	"sync/atomic"
)

func DefaultPickStrategy[T any](circularIdx int64) func(slice []*Limiter[T]) *Limiter[T] {
	return func(slice []*Limiter[T]) *Limiter[T] {
		atomic.AddInt64(&circularIdx, 1)

		if circularIdx >= int64(len(slice)) {
			circularIdx = 0
		}

		return slice[circularIdx]
	}
}
