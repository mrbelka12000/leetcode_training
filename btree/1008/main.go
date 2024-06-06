package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	t := &TreeNode{
		Val: preorder[0],
	}

	for i := 1; i < len(preorder); i++ {
		Insert(t, preorder[i])
	}

	return t
}

func Insert(t *TreeNode, val int) {
	if t == nil {
		t = &TreeNode{
			Val: val,
		}
		return
	}
	if t.Val >= val {
		if t.Left == nil {
			t.Left = &TreeNode{
				Val: val,
			}
		} else {
			Insert(t.Left, val)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{
				Val: val,
			}
		} else {
			Insert(t.Right, val)
		}
	}
}
