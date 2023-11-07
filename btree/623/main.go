package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		tmp := &TreeNode{
			Val:  val,
			Left: root,
		}

		return tmp
	}
	addRows(root, val, depth, 1)
	return root
}

func addRows(root *TreeNode, val, dep, currDep int) {
	if root == nil {
		return
	}

	if dep-1 == currDep {
		lTmp := &TreeNode{
			Val:  val,
			Left: root.Left,
		}

		rTmp := &TreeNode{
			Val:   val,
			Right: root.Right,
		}

		root.Left = lTmp
		root.Right = rTmp
		return
	}

	addRows(root.Left, val, dep, currDep+1)
	addRows(root.Right, val, dep, currDep+1)
}
