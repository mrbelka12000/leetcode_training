package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	return getRight(root)
}

func getRight(root *Node) *Node {
	q := []*Node{root}

	for len(q) != 0 {

		size := len(q)

		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]

			if i < size-1 {
				// fmt.Println(n.Val, q, i, size)
				n.Next = q[0]
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		// fmt.Println()
	}

	return root
}
