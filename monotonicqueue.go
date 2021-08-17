package godtype

import (
	"reflect"
)

type MonotonicQueue struct {
	deque	*Deque
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{NewDeque()}
}

func (mq *MonotonicQueue) Pop(elem interface{}) {
	if elem == mq.deque.PeekFirst() {
		mq.deque.PopFirst()
	}
}

func (mq *MonotonicQueue) Push(elem interface{}) {
	switch v := reflect.ValueOf(elem); v.Kind() {
	case reflect.Int:
		for ! mq.deque.IsEmpty() && mq.deque.PeekLast().(int) < int(v.Int()) {
			mq.deque.PopLast()
		}
	case reflect.Float64:
		for ! mq.deque.IsEmpty() && mq.deque.PeekLast().(float64) < v.Float() {
			mq.deque.PopLast()
		}
	case reflect.String:
		for ! mq.deque.IsEmpty() && mq.deque.PeekLast().(string) < v.String() {
			mq.deque.PopLast()
		}
	default:
		return
	}

	mq.deque.PushLast(elem)
}

func (mq *MonotonicQueue) Max() interface{} {
	return mq.deque.PeekFirst()
}

func (mq *MonotonicQueue) IsEmpty() bool {
	return mq.deque.IsEmpty()
}

func (mq *MonotonicQueue) Size() int {
	return mq.deque.Size()
}
