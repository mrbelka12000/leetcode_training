package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	left := []*TreeNode{}
	right := []*TreeNode{}

	findPath(root, p, &left)
	findPath(root, q, &right)

	mp := make(map[int]*TreeNode)

	for i := len(left) - 1; i >= 0; i-- {
		// fmt.Println(left[i].Val)
		mp[left[i].Val] = left[i]
	}

	for i := len(right) - 1; i >= 0; i-- {
		if val, ok := mp[right[i].Val]; ok {
			return val
		}
	}

	return root
}

func findPath(start, end *TreeNode, result *[]*TreeNode) bool {
	if start == nil {
		return false
	}

	if start.Val == end.Val {
		return true
	}

	if findPath(start.Left, end, result) {
		*result = append([]*TreeNode{start.Left}, *result...)
		return true
	}
	if findPath(start.Right, end, result) {
		*result = append([]*TreeNode{start.Right}, *result...)
		return true
	}

	return false
}
