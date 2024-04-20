package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root == nil {
		return 0
	}

	q := []*Node{root}
	var result int

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			n := q[0]
			q = q[1:]
			for _, v := range n.Children {
				q = append(q, v)
			}
		}
		result++
	}

	return result
}
