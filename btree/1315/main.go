package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	var sum int
	getChilds(nil, root, root.Left, &sum)
	getChilds(nil, root, root.Right, &sum)

	return sum
}

func getChilds(gp, parr, curr *TreeNode, sum *int) {
	if parr == nil || curr == nil {
		return
	}
	if gp != nil {
		if gp.Val%2 == 0 {
			*sum += curr.Val
		}
	}

	getChilds(parr, curr, curr.Left, sum)
	getChilds(parr, curr, curr.Right, sum)
}
