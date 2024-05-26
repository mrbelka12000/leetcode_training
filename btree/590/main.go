package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	var res []int

	for _, v := range root.Children {
		res = append(res, postorder(v)...)
	}

	res = append(res, root.Val)

	return res
}
