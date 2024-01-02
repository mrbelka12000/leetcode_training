package main

func main() {

}

func findWords(board [][]byte, words []string) []string {
	t := NewTrie()

	for _, word := range words {
		t.AddWord(word)
	}

	var result []string

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if val, ok := t.n.childs[board[i][j]]; ok {
				t.runner(board, i, j, []byte{board[i][j]}, val, &result)
			}
		}
	}
	return result
}

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (t *Trie) runner(board [][]byte, x, y int, tmp []byte, curr *Node, result *[]string) {
	if curr != nil && curr.end {
		*result = append(*result, string(tmp))
		curr.val--
		if curr.val == 0 {
			curr.end = false
		}
	}

	prev := board[x][y]
	board[x][y] = 1

	for _, dir := range dirs {
		xx, yy := x+dir[0], y+dir[1]
		if xx < 0 || xx >= len(board) || yy < 0 || yy >= len(board[xx]) || board[xx][yy] == 1 {
			continue
		}

		val, ok := curr.childs[board[xx][yy]]
		if !ok {
			continue
		}

		t.runner(board, xx, yy, append(tmp, board[xx][yy]), val, result)
	}

	board[x][y] = prev
}

type (
	Trie struct {
		n *Node
	}
	Node struct {
		end    bool
		val    int
		childs map[byte]*Node
	}
)

func NewTrie() *Trie {
	return &Trie{
		n: NewNode(),
	}
}

func NewNode() *Node {
	return &Node{
		childs: make(map[byte]*Node),
	}
}

func (t *Trie) AddWord(word string) {
	curr := t.n

	for i := 0; i < len(word); i++ {
		_, ok := curr.childs[word[i]]
		if !ok {
			curr.childs[word[i]] = NewNode()
		}

		curr = curr.childs[word[i]]
	}

	curr.end = true
	curr.val++
}

/*
[["o","a","a","n"]
,["e","t","a","e"]
,["i","h","k","r"]
,["i","f","l","v"]]
["oath","pea","eat","rain","oathi","oathk","oathf","oate","oathii","oathfi","oathfii"]


["oate","oath","eat"]
["oath","oathk","oathf","oathfi","oathfii","oathi","oathii","oate","eat"]
*/
