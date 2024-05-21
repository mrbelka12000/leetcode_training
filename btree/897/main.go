package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	tmp := &TreeNode{}
	t := tmp
	if root == nil {
		return nil
	}

	var runner func(check *TreeNode)
	runner = func(check *TreeNode) {
		if check == nil {
			return
		}
		runner(check.Left)
		t.Right = check
		t = t.Right
		check.Left = nil
		runner(check.Right)
	}

	runner(root)
	return tmp.Right
}
