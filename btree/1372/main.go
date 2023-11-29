package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	var result int

	dfs(root, &result)

	return result
}

func dfs(root *TreeNode, maxNum *int) {
	if root == nil {
		return
	}
	var l, r int = -1, -1

	getPath(root, true, &l)
	// fmt.Println("====================")
	getPath(root, false, &r)

	*maxNum = max(*maxNum, max(l, r))
	// fmt.Println(l,r)
	dfs(root.Left, maxNum)
	dfs(root.Right, maxNum)
}

func getPath(root *TreeNode, dir bool, count *int) {
	if root == nil {
		return
	}

	// fmt.Println(dir)

	*count++

	if dir {
		getPath(root.Left, !dir, count)
	} else {
		getPath(root.Right, !dir, count)
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
