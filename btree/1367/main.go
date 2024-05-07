package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	return runner(head, root)
}

func runner(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	if root.Val == head.Val {
		val := dfs(head.Next, root.Left) || dfs(head.Next, root.Right)

		if val {
			return true
		}
	}

	return runner(head, root.Left) ||
		runner(head, root.Right)
}

func dfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Val != head.Val {
		return false
	}

	return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
}
