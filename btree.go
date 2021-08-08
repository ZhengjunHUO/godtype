package godtype

import (
	"fmt"
)

type TreeNode struct {
     Val	interface{}
     Left 	*TreeNode
     Right 	*TreeNode
}

// 从数列(完全二叉树标准)构建树
func NewBTree(elems []interface{}, index int) *TreeNode {
        // 为叶子结点
        if 2*index + 2 >= len(elems) {
		return &TreeNode{ elems[index], nil, nil, }
	}

        // 有子节点
	var l, r *TreeNode
        if 2*index + 1 < len(elems) && elems[2*index+1] != nil {
		l = NewBTree(elems, 2*index+1)
	}
	
        if 2*index + 2 < len(elems) && elems[2*index+2] != nil {
		r = NewBTree(elems, 2*index+2)
	}
	
	return &TreeNode{
		Val: elems[index],
		Left: l,
		Right: r,
        }
}

func PrintBtree(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty tree")
		return
	}

	fmt.Println("Current node's value: ", root.Val)
	if root.Left !=	nil {
		fmt.Printf("%v have a left child\n", root.Val)
		PrintBtree(root.Left)
	}
	if root.Right != nil {
		fmt.Printf("%v have a right child\n", root.Val)
		PrintBtree(root.Right)
	}
}

func PrintBtreeBFS(root *TreeNode) {
	rslt := []interface{}{}
	if root == nil {
		fmt.Println(rslt)
	}

	q := NewQueue()
	q.Push(root)

	var emptyNode TreeNode
	size := 0

	loop: for !q.IsEmpty() {
		size = q.Size()
		emptyNum := 0
		for i:=0; i<size; i++ {
			node := q.Pop().(*TreeNode)
			if *node != emptyNode {
				rslt = append(rslt, node.Val)
			}else{
				emptyNum++
				rslt = append(rslt, nil)
			}

			if emptyNum == size {
				break loop
			}

			if node.Left != nil {
				q.Push(node.Left)
			}else{
				q.Push(&TreeNode{})
			}

			if node.Right != nil {
				q.Push(node.Right)
			}else{
				q.Push(&TreeNode{})
			}

		}
	}

	fmt.Println(rslt[:len(rslt)-size])
}
