package godtype

import (
	"fmt"
)

type ListNode struct {
        Val int
        Next *ListNode
}

func NewList(l []int) *ListNode {
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

func PrintList(node *ListNode) {
        curr := node
        for {
		fmt.Println(curr.Val)

                if curr.Next == nil {
                        break
                }
                curr = curr.Next
        }
}
