package main

func main() {

}

type (
	BrowserHistory struct {
		List *Node
	}

	Node struct {
		Val  string
		Next *Node
		Prev *Node
	}
)

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		List: &Node{
			Val: homepage,
		},
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.List.Next = &Node{
		Val:  url,
		Prev: this.List,
	}

	this.List = this.List.Next
}

func (this *BrowserHistory) Back(steps int) string {
	for this.List.Prev != nil {
		if steps == 0 {
			break
		}
		steps--
		this.List = this.List.Prev
	}

	return this.List.Val
}

func (this *BrowserHistory) Forward(steps int) string {
	for this.List.Next != nil {
		if steps == 0 {
			break
		}
		steps--
		this.List = this.List.Next
	}
	return this.List.Val
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
