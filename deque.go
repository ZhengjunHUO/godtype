package godtype

import (
	"sync"
)

type Deque struct {
	Elems []interface{}
	Lock  sync.RWMutex
}

func NewDeque() *Deque {
	return &Deque{}
}

func (d *Deque) PushFirst(elem interface{}) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elems = append([]interface{}{elem}, d.Elems...)
}

func (d *Deque) PushLast(elem interface{}) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	d.Elems = append(d.Elems, elem)
}

func (d *Deque) PopFirst() interface{} {
	if d.IsEmpty() {
		return nil
	}

	d.Lock.Lock()
	defer d.Lock.Unlock()

	defer func(){
		d.Elems[0] = nil
		d.Elems = d.Elems[1:]
	}()
	
	return d.Elems[0]
}

func (d *Deque) PopLast() interface{} {
	if d.IsEmpty() {
		return nil
	}

	d.Lock.Lock()
	defer d.Lock.Unlock()

	n := len(d.Elems)

	defer func(){
		d.Elems[n-1] = nil
		d.Elems = d.Elems[:n-1]
	}()
	
	return d.Elems[n-1]
}

func (d *Deque) PeekFirst() interface{} {
	if d.IsEmpty() {
		return nil
	}

	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elems[0]
}

func (d *Deque) PeekLast() interface{} {
	if d.IsEmpty() {
		return nil
	}

	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return d.Elems[len(d.Elems)-1]
}

func (d *Deque) IsEmpty() bool {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elems) == 0 
}

func (d *Deque) Size() int {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	return len(d.Elems)
}
