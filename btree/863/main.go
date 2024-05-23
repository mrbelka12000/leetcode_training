package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	result = nil

	runner(root, target, k)

	return result
}

func runner(root, target *TreeNode, k int) (int, bool) {
	if root == nil {
		return -1, false
	}

	if target == root {
		runnerChild(root, k)
		return k - 1, true
	}

	newDist, ok := runner(root.Left, target, k)
	if ok {
		if newDist == 0 {
			result = append(result, root.Val)
			return -1, false
		}
		runnerChild(root.Right, newDist-1)
		return newDist - 1, true
	}

	newDist, ok = runner(root.Right, target, k)
	if ok {
		if newDist == 0 {
			result = append(result, root.Val)
			return -1, false
		}
		runnerChild(root.Left, newDist-1)
		return newDist - 1, true
	}

	return -1, false
}

var result []int

func runnerChild(root *TreeNode, k int) {
	if root == nil || k < 0 {
		return
	}

	if k == 0 {
		result = append(result, root.Val)
		return
	}

	runnerChild(root.Left, k-1)
	runnerChild(root.Right, k-1)
}
