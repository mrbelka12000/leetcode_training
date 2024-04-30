package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	return runner(root)
}

func runner(root *TreeNode) bool {
	var check bool
	q := []*TreeNode{root}

	for len(q) != 0 {

		size := len(q)

		prev := -1
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			// fmt.Println(n.Val % 2, check)
			if !check {
				if n.Val != 1 && n.Val%2 != 1 {
					return false
				}
			} else {
				if n.Val%2 != 0 {
					return false
				}
			}

			if prev != -1 {
				if check {
					if prev <= n.Val {
						return false
					}
				} else {
					if prev >= n.Val {
						return false
					}
				}
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
			prev = n.Val
		}
		// fmt.Println()
		check = !check
	}
	return true
}
