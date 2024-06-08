package main

func main() {

}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {
	if getLCA(p, q) {
		return p
	}
	if getLCA(q, p) {
		return q
	}

	m = make(map[*Node]bool)

	return getParentPath(p, q)
}

var m = map[*Node]bool{}

func getParentPath(p, q *Node) *Node {
	if p == nil && q == nil {
		return nil
	}
	if m[p] {
		return p
	}
	if p != nil {
		m[p] = true
	}

	if m[q] {
		return q
	}
	if q != nil {
		m[q] = true
	}

	if p == nil {
		return getParentPath(p, q.Parent)
	}
	if q == nil {
		return getParentPath(p.Parent, q)
	}

	return getParentPath(p.Parent, q.Parent)
}

func getLCA(f, s *Node) bool {
	if f == nil {
		return false
	}

	return f == s || getLCA(f.Left, s) || getLCA(f.Right, s)
}
