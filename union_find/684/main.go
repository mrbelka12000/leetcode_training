package main

func main() {

}

func findRedundantConnection(edges [][]int) []int {
	s := &Solution{}

	return s.findRedundantConnection(edges)
}

type Solution struct {
	parent []int
}

func (s *Solution) makeSet(mp map[int]int) {
	s.parent = make([]int, len(mp)+1)
	for k := range mp {
		s.parent[k] = k
	}
}

func (s *Solution) Find(x int) int {
	if s.parent[x] != x {
		s.parent[x] = s.Find(s.parent[x])
	}
	return s.parent[x]
}

func (s *Solution) Union(x int, y int) {
	xSet := s.Find(x)
	ySet := s.Find(y)

	if xSet == ySet {
		return
	}

	s.parent[ySet] = xSet
}

func (s *Solution) findRedundantConnection(edges [][]int) []int {
	// ToDo: Write Your Code Here.
	mp := make(map[int]int)
	for _, v := range edges {
		mp[v[0]]++
		mp[v[1]]++
	}

	// fmt.Println()
	s.makeSet(mp)

	for _, v := range edges {
		x, y := s.Find(v[0]), s.Find(v[1])
		if x == y {
			return []int{v[0], v[1]}
		}
		s.Union(v[0], v[1])
	}

	return []int{0, 0} // If no redundant connection found, though the problem statement assures there will be one.
}
