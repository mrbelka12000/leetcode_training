package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	q []int
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{}

	bst.getNode(root)

	return bst
}

func (this *BSTIterator) getNode(root *TreeNode) {
	if root == nil {
		return
	}

	this.getNode(root.Left)
	this.q = append(this.q, root.Val)
	this.getNode(root.Right)
}

func (this *BSTIterator) Next() int {
	val := this.q[0]

	this.q = this.q[1:]

	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.q) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
