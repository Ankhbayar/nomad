package stats

import (
	"fmt"
	"sync"
)

var (
	// The default size of the ring buffer
	defaultCap = 60
)

// RingBuff is a data structure which is a circular list based on slices
type RingBuff struct {
	head int
	buff []interface{}

	lock sync.RWMutex
}

// NewRingBuff creates a new ring buffer of the specified size
func NewRingBuff(capacity int) (*RingBuff, error) {
	if capacity < 1 {
		return nil, fmt.Errorf("can not create a ring buffer with capacity: %v", capacity)
	}
	return &RingBuff{buff: make([]interface{}, 0, capacity), head: -1}, nil
}

// Enqueue queues a new value in the ring buffer. This operation would
// over-write an older value if the list has reached it's capacity
func (r *RingBuff) Enqueue(value interface{}) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if len(r.buff) < cap(r.buff) {
		r.buff = append(r.buff, struct{}{})
	}
	r.head += 1
	if r.head == cap(r.buff) {
		r.head = 0
	}
	r.buff[r.head] = value
}

// Peek returns the last value enqueued in the ring buffer
func (r *RingBuff) Peek() interface{} {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if r.head == -1 {
		return nil
	}
	return r.buff[r.head]
}

// Values returns all the values in the buffer.
func (r *RingBuff) Values() []interface{} {
	r.lock.RLock()
	defer r.lock.RUnlock()
	if r.head == len(r.buff)-1 {
		vals := make([]interface{}, len(r.buff))
		copy(vals, r.buff)
		return vals
	}

	slice1 := r.buff[r.head+1:]
	slice2 := r.buff[:r.head+1]
	return append(slice1, slice2...)
}
