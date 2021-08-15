package godtype

// Doubly Linked Node
type DoublyLNode struct {
	Key, Val	interface{}	
	Prev, Next	*DoublyLNode
}

func NewDoublyLNode(k, v interface{}) *DoublyLNode {
	return &DoublyLNode{k, v, nil, nil}
}


// Doubly Linked List
type DoublyLList struct {
	Head, Tail	*DoublyLNode
	Len		int 
}

func NewDoublyLList() *DoublyLList {
	Head, Tail := NewDoublyLNode(nil, nil), NewDoublyLNode(nil, nil)
	Head.Next, Tail.Prev = Tail, Head

	return &DoublyLList{Head, Tail, 0}
}

// Push at the end
func (dl *DoublyLList) Push(node *DoublyLNode) {
	node.Prev, node.Next = dl.Tail.Prev, dl.Tail
	dl.Tail.Prev.Next, dl.Tail.Prev = node, node	
	dl.Len++
} 

// Pop from the beginning
func (dl *DoublyLList) Pop() *DoublyLNode {
	if dl.Head.Next == dl.Tail {
		return nil
	}

	target := dl.Head.Next
	dl.Delete(target)
	return target
}

func (dl *DoublyLList) Delete(node *DoublyLNode) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	dl.Len--
}


// Linked Hash Map
type Lmap struct {
	Map	map[interface{}]*DoublyLNode
	DList	*DoublyLList
}

func NewLmap() *Lmap {
	return &Lmap{make(map[interface{}]*DoublyLNode), NewDoublyLList()}
}

func (lm *Lmap) Put(k, v interface{}) {
	if _, ok := lm.Map[k]; ok {
		lm.Map[k].Val = v
		return
	}
		
	node := NewDoublyLNode(k, v)
	lm.DList.Push(node)
	lm.Map[k] = node
}

func (lm *Lmap) Get(k interface{}) interface{} {
	if _, ok := lm.Map[k]; !ok {
		return nil
	}

	return lm.Map[k].Val
}

func (lm *Lmap) Delete(k interface{}) {
	if _, ok := lm.Map[k]; !ok {
		return
	}

	lm.DList.Delete(lm.Map[k])
	delete(lm.Map, k)
}

func (lm *Lmap) PopEldest() *DoublyLNode {
	node := lm.DList.Pop()
	delete(lm.Map, node.Key)
	return node
}

func (lm *Lmap) BecomeNewest(k interface{}) {
	if _, ok := lm.Map[k]; !ok {
		return
	}
	
	node := lm.Map[k]
	lm.DList.Delete(node)
	lm.DList.Push(node)
}

func (lm *Lmap) Contains(k interface{}) bool {
	if _, ok := lm.Map[k]; !ok {
		return false
	}

	return true
}

func (lm *Lmap) Size() int {
	return lm.DList.Len
}
