package godtype

import (
	"container/heap"
	"sync"
	"reflect"
	"fmt"
)

type Elem struct {
	Value		interface{}
	Priority	interface{}
	Index		int
}

type Heap []*Elem

type PriorityQueue struct {
	Data		Heap
	Lock		sync.RWMutex
	PopLowest	bool
}

// build a priority queue with a slice of value and a slice of priority
func NewPQ(values interface{}, prios interface{}, popLowest bool) *PriorityQueue {
	v := reflect.ValueOf(values)
        if v.Kind() != reflect.Slice {
                fmt.Println("Input values is not a slice !")
                return nil
        }

	p := reflect.ValueOf(prios)
        if v.Kind() != reflect.Slice {
                fmt.Println("Input prios is not a slice !")
                return nil
        }

	nv := reflect.ValueOf(values).Len()
	np := reflect.ValueOf(prios).Len()
	if nv != np {
                fmt.Println("Length of values and prios doesn't match !")
                return nil
	}

	sign := 1
	/* 
	  according to Less() the pq will pop the value lowest priority
	  if we want the opposite behavior, do *(-1)
	*/
	if !popLowest {
		sign = -1
	}

	data := make(Heap, nv)
	for i:=0; i<nv; i++ {
		var pr interface{}
		if p.Index(i).Kind() == reflect.Float64 {
			pr = p.Index(i).Float() * float64(sign)
		}else{
			pr = int(p.Index(i).Int()) * sign
		}

		data[i] = &Elem{
			Value:		v.Index(i).Interface(),
			Priority:	pr,
			Index:		i,
		}
	}
	heap.Init(&data)

	return &PriorityQueue{
		Data: data,
		PopLowest: popLowest,
	}
}

// Implement sort interface
func (h Heap) Len() int {
	return len(h)
}

// Implement sort interface
func (h Heap) Less(i, j int) bool {
	v := reflect.ValueOf(h[i].Priority)
	if v.Kind() == reflect.Float64 {
		return h[i].Priority.(float64) < h[j].Priority.(float64)
	}

	return h[i].Priority.(int) < h[j].Priority.(int)
}

// Implement sort interface
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

// Implement heap interface
func (h *Heap) Push(item interface{}) {
	elem := item.(*Elem)
	elem.Index = h.Len()
	*h = append(*h, elem)
}

// Implement heap interface
func (h *Heap) Pop() interface{} {
	n := h.Len()

	defer func() {
		(*h)[n-1] = nil
		*h = (*h)[:n-1]
	}()

	return (*h)[n-1]
}

// Update priority of element, if elem not exist insert it
/*
func (pq *PriorityQueue) Update(value interface{}, prio int) {
	pq.Lock.Lock()
	defer pq.Lock.Unlock()

	for _, elem := range pq.Data {
		if elem.Value == value {
			if pq.PopLowest {
				elem.Priority = prio
			}else{
				elem.Priority = prio * (-1)
			}
			heap.Fix(&(pq.Data), elem.Index)
			return
		}
	}
	
	pq.Push(value, prio)
}
*/

func (pq *PriorityQueue) Remove(value interface{}) interface{} {
	pq.Lock.Lock()
	defer pq.Lock.Unlock()

	for _, elem := range pq.Data {
		if elem.Value == value {
			return heap.Remove(&(pq.Data), elem.Index)
		}
	}

	return nil
}

func (pq *PriorityQueue) Pop() interface{} {
	if pq.Data.Len() < 1 {
		return nil
	}

	pq.Lock.Lock()
	defer pq.Lock.Unlock()

	return heap.Pop(&(pq.Data)).(*Elem).Value
}

func (pq *PriorityQueue) PopWithPrio() []interface{} {
        if pq.Data.Len() < 1 {
                return nil
        }

        pq.Lock.Lock()
        defer pq.Lock.Unlock()

	sign := 1
	if !pq.PopLowest {
                sign = -1
        }

	elem := heap.Pop(&(pq.Data)).(*Elem)

        if reflect.ValueOf(elem.Priority).Kind() == reflect.Float64 {
                return []interface{}{elem.Value, elem.Priority.(float64) * float64(sign)}
        }

        return []interface{}{elem.Value, elem.Priority.(int) * sign}
}

func (pq *PriorityQueue) Push(value interface{}, prio interface{}) {
	pq.Lock.Lock()
	defer pq.Lock.Unlock()

	sign := 1
	if !pq.PopLowest {
		sign = -1
	}

	var pr interface{}
	if reflect.ValueOf(prio).Kind() == reflect.Float64 {
		pr = prio.(float64) * float64(sign)
	}else{
		pr = prio.(int) * sign
	}

	elem := &Elem{
		Value:		value,
		Priority:	pr,
	}
	heap.Push(&(pq.Data), elem)
}

// Pop without remove elem from PQ
func (pq *PriorityQueue) Peek() interface{} {
	if pq.Data.Len() < 1 {
		return nil
	}

	pq.Lock.RLock()
	defer pq.Lock.RUnlock()

	return pq.Data[0].Value
}

func (pq *PriorityQueue) Size() int {
	return pq.Data.Len()
}
