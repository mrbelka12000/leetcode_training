package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	mp := make(map[int]bool)

	for _, v := range to_delete {
		mp[v] = true
	}

	var result []*TreeNode

	runner(root, nil, mp, &result, true, true)

	return result
}

func runner(root, parent *TreeNode, mp map[int]bool, result *[]*TreeNode, deleted, dir bool) {
	if root == nil {
		return
	}

	if mp[root.Val] {
		runner(root.Left, root, mp, result, true, false)
		runner(root.Right, root, mp, result, true, true)
		if parent != nil {
			if dir {
				parent.Right = nil
			} else {
				parent.Left = nil
			}
		}
		return
	}

	if deleted {
		*result = append(*result, root)
	}

	runner(root.Left, root, mp, result, false, false)
	runner(root.Right, root, mp, result, false, true)
}
