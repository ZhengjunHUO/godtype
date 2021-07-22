package godtype

import (
	"container/heap"
	"reflect"
	"fmt"
)

// Decide how the Less() compare the priority
var PopLowest bool

type Elem struct {
	Value		interface{}
	Priority	int
	Index		int
}

type PriorityQueue []*Elem

// build a priority queue with a slice of value and a slice of priority
func InitPQ(values interface{}, prios []int, popLowest bool) PriorityQueue {
	PopLowest = popLowest

	v := reflect.ValueOf(values)
        if v.Kind() != reflect.Slice {
                fmt.Println("Input values is not a slice !")
                return nil       
        }

	n := reflect.ValueOf(values).Len()
	if n != len(prios) {
                fmt.Println("Length of values and prios doesn't match !")
                return nil       
	}

	pq := make(PriorityQueue, n)
	for i:=0; i<n; i++ {
		pq[i] = &Elem{
			Value:		v.Index(i).Interface(),
			Priority:	prios[i],
			Index:		i,
		}
	} 
	heap.Init(&pq)

	return pq
}

// Implement sort interface
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Implement sort interface
func (pq PriorityQueue) Less(i, j int) bool {
	if PopLowest {
		// Will pop the lowest priority value first
		return pq[i].Priority < pq[j].Priority
	}

	// Will pop the highest priority value first
	return pq[i].Priority > pq[j].Priority
}

// Implement sort interface
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Implement heap interface
func (pq *PriorityQueue) Push(item interface{}) {
	elem := item.(*Elem)
	elem.Index = pq.Len()
	*pq = append(*pq, elem)
}

// Implement heap interface
func (pq *PriorityQueue) Pop() interface{} {
	n := pq.Len()

	defer func() {
		(*pq)[n-1] = nil
		*pq = (*pq)[:n-1]
	}()

	return (*pq)[n-1]	
}

// Update priority of element, if elem not exist insert it
func (pq *PriorityQueue) Update(value interface{}, prio int) {
	for _, elem := range *pq {
		if elem.Value == value {
			elem.Priority = prio
			heap.Fix(pq, elem.Index)
			return
		}
	}
	
	pq.Insert(value, prio)
}

// Pop
func (pq *PriorityQueue) Pull() interface{} {
	return heap.Pop(pq).(*Elem).Value
}

// Push
func (pq *PriorityQueue) Insert(value interface{}, prio int) {
	elem := &Elem{
		Value:		value,
		Priority:	prio,
	}
	heap.Push(pq, elem)
}

// Pop without remove elem from PQ
func (pq *PriorityQueue) Peek() interface{} {
	return (*pq)[0].Value
}
