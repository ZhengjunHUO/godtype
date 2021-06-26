package godtype

import (
	"fmt"
	"reflect"
)

type ListNode struct {
        Val interface{}
        Next *ListNode
}

func NewList(l interface{}) *ListNode {
	v := reflect.ValueOf(l)
	if v.Kind() != reflect.Slice {
		fmt.Println("Input is not a slice !")
		return nil	 
	}

	n := reflect.ValueOf(l).Len()
	if n == 0 {
		return nil
	} 

        head := &ListNode{v.Index(0).Interface(), nil}
        if n == 1 {
                return head
        }

        curr := head
        for i:=1; i<n; i++ {
                curr.Next = &ListNode{v.Index(i).Interface(), nil}
                curr = curr.Next
        }

        return head
}

func PrintList(node *ListNode) {
	if node != nil {
	        curr := node
	        for {
			fmt.Printf("%v, ", curr.Val)

	                if curr.Next == nil {
				fmt.Println()
	                        break
	                }
	                curr = curr.Next
	        }
	}
}
