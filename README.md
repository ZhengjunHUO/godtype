# godtype

This is the homemade thread-safe Go data type mentioned in my [leetcode project](https://github.com/ZhengjunHUO/leetcode), initally built to solve some leetcode problems and turns out to be quite useful in other places.

## Usage
1. \# go get [-u] github.com/ZhengjunHUO/godtype

2. Create a new struct:        
  - Stack: godtype.NewStack()
  - Queue: godtype.NewQueue()
  - Deque: godtype.NewDeque()
  
3. Method:


| Operation | [Stack](https://github.com/ZhengjunHUO/godtype/blob/main/stack.go) | [Queue](https://github.com/ZhengjunHUO/godtype/blob/main/queue.go) | [Deque](https://github.com/ZhengjunHUO/godtype/blob/main/deque.go) |
|----- | ----- | ----- | ----- |
| insert at head | - | - | PushFirst(elem interface{}) | 
| insert at end | Push(elem interface{}) | Push(elem interface{}) | PushLast(elem interface{}) |
| remove from head | - | Pop() interface{} | PopFirst() interface{} |
| remove from end | Pop() interface{} | - | PopLast() interface{} |
| check the head | - | - | PeekFirst() interface{} | 
| check the end | Peek() interface{} | - | PeekLast() interface{} |
| check is empty | IsEmpty() bool | IsEmpty() bool | IsEmpty() bool |
| check size | Size() int | Size() int | Size() int |

### Notice:

The element is interface{} type, do a type assertion if needed after a get method (eg. dq.PeekFirst().(int) > 0).
