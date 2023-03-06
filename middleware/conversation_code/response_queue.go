package conversation_code

import (
	"errors"
	"sync"
	humanize_protobuf "testserver/src/generated/humanize-protobuf"
)

type ResponseQueueImpl struct {
	mu       sync.Mutex
	capacity int
	q        []*humanize_protobuf.MessageResponseData
}

type ResponseQueue interface {
	Insert(item *humanize_protobuf.MessageResponseData) error
	Remove() (*humanize_protobuf.MessageResponseData, error)
	IsEmpty() bool
}

// CreateResponseQueue creates an empty ResponseQueueImpl with desired capacity
func CreateResponseQueue(capacity int) ResponseQueue {
	return &ResponseQueueImpl{
		capacity: capacity,
		q:        make([]*humanize_protobuf.MessageResponseData, 0, capacity),
	}
}

// Insert inserts the item into the ResponseQueueImpl
func (q *ResponseQueueImpl) Insert(item *humanize_protobuf.MessageResponseData) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) < int(q.capacity) {
		q.q = append(q.q, item)
		return nil
	}
	return errors.New("ResponseQueueImpl is full")
}

// Remove removes the oldest element from the ResponseQueueImpl
func (q *ResponseQueueImpl) Remove() (*humanize_protobuf.MessageResponseData, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) > 0 {
		item := q.q[0]
		q.q = q.q[1:]
		return item, nil
	}
	return nil, errors.New("ResponseQueueImpl is empty")
}

func (q *ResponseQueueImpl) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.q) == 0 {
		return true
	}
	return false
}
