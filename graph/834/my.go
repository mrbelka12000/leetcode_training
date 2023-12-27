package main

func main() {

}

func sumOfDistancesInTreeMy(n int, edges [][]int) []int {
	g := &G{}

	for _, v := range edges {
		g.AddEdje(v[0], v[1])
		g.AddEdje(v[1], v[0])
	}

	if len(g.nodes) == 0 {
		return []int{0}
	}

	result := make([]int, n)
	for i := 0; i < len(result); i++ {
		result[i] = g.CalculateDistance(i)
	}

	return result
}

type (
	G struct {
		nodes []*node
	}
	node struct {
		val   int
		edjes []*node
	}
)

func (g *G) AddNode(val int) *node {
	n := &node{val: val}

	g.nodes = append(g.nodes, n)

	return n
}

func (g *G) GetNode(val int) *node {
	for _, v := range g.nodes {
		if v.val == val {
			return v
		}
	}
	return nil
}

func (g *G) AddEdje(from, to int) {
	fromNode, toNode := g.GetNode(from), g.GetNode(to)
	if fromNode == nil {
		fromNode = g.AddNode(from)
	}
	if toNode == nil {
		toNode = g.AddNode(to)
	}

	fromNode.edjes = append(fromNode.edjes, toNode)
}

func (g *G) CalculateDistance(val int) int {
	n := g.GetNode(val)

	mp := make(map[int]int)
	visited := make(map[int]struct{})

	q := []*node{n}

	var dist int
	for len(q) != 0 {
		dist++

		size := len(q)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			visited[node.val] = struct{}{}
			for _, v := range node.edjes {
				if _, ok := visited[v.val]; ok {
					continue
				}
				mp[v.val] = dist
				q = append(q, v)
			}
		}
	}

	var result int

	for _, v := range mp {
		result += v
	}

	return result
}
