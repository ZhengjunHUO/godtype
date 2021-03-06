package godtype

import (
	"sync"
)

type Queue struct {
	Elems []interface{}
	Lock  sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{}
}

// Push to the queue's end
func (q *Queue) Push(elem interface{}) {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	q.Elems = append(q.Elems, elem)
}

// Pop from the queue's head
func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}

	q.Lock.Lock()
	defer q.Lock.Unlock()

	defer func(){
		q.Elems[0] = nil
		q.Elems = q.Elems[1:]
	}()
	
	return q.Elems[0]
}

func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	q.Lock.RLock()
	defer q.Lock.RUnlock()

	return q.Elems[0]
}

func (q *Queue) IsEmpty() bool {
	q.Lock.RLock()
	defer q.Lock.RUnlock()

	return len(q.Elems) == 0 
}

func (q *Queue) Size() int {
	q.Lock.RLock()
	defer q.Lock.RUnlock()

	return len(q.Elems)
}
