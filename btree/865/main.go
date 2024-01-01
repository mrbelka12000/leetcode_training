package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	node, _ := lca(root)
	return node
}

func lca(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	left, dl := lca(root.Left)
	right, dr := lca(root.Right)

	if dl == dr {
		return root, dl + 1
	} else if dl > dr {
		return left, dl + 1
	}
	return right, dr + 1
}
