package godtype

type ListNode struct {
        Val interface{}
        Next *ListNode
}

func NewList(l []interface{}) *ListNode {
        n := len(l)
	if n == 0 {
		return nil
	} 

        head := &ListNode{l[0], nil}
        if n == 1 {
                return head
        }

        curr := head
        for i:=1; i<n; i++ {
                curr.Next = &ListNode{l[i], nil}
                curr = curr.Next
        }

        return head
}
