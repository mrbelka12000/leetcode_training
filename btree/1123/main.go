package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	p, q := bfs(root)

	// fmt.Println(p, q)
	return lca(root, p, q)
}

func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lca(root.Left, p, q)
	right := lca(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if right != nil {
		return right
	}

	return left
}

func bfs(root *TreeNode) (*TreeNode, *TreeNode) {
	q := []*TreeNode{root}
	var result []*TreeNode
	for len(q) != 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			n := q[0]
			result = append(result, n)
			q = q[1:]

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		if len(q) != 0 {
			result = make([]*TreeNode, 0)
		}
	}

	if len(result) == 1 {
		return result[0], result[0]
	}

	return result[0], result[len(result)-1]
}
