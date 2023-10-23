package main

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
	q := []*Node{root}

	for len(q) != 0 {
		count := len(q)
		for i := 0; i < count; i++ {
			n := q[0]
			q = q[1:]
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			if len(q) == 0 || count-i == 1 {
				break
			}

			n.Next = q[0]
		}
	}

	return root
}
