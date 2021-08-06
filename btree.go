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
		l = NewTree(elems, 2*index+1)
	}
	
        if 2*index + 2 < len(elems) && elems[2*index+2] != nil {
		r = NewTree(elems, 2*index+2)
	}
	
	return &TreeNode{
		Val: elems[index],
		Left: l,
		Right: r,
        }
}

func PrintBtree(root *TreeNode) {
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
