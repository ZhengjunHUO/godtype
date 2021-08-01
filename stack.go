package godtype

import (
	"sync"
)

type Stack struct {
	Elems	[]interface{}
	Lock	sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return len(s.Elems) == 0
}

func (s *Stack) Push(elem interface{}) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Elems = append(s.Elems, elem)	
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}	

	s.Lock.Lock()
	defer s.Lock.Unlock()

	n := len(s.Elems)

	defer func(){
		s.Elems[n-1] = nil
		s.Elems = s.Elems[:n-1]
	}() 

	return s.Elems[n-1] 
}

func (s *Stack) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return len(s.Elems)
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}	

	s.Lock.RLock()
	defer s.Lock.RUnlock()

	return s.Elems[len(s.Elems)-1]
}
