package main

import (
	"fmt"
	"sort"
)

func main() {
	t := Constructor()

	t.PostTweet(1, 5)
	t.Follow(1, 2)
	t.Follow(2, 1)
	fmt.Println(t.GetNewsFeed(2))

	t.PostTweet(2, 6)

	fmt.Println(t.GetNewsFeed(1))
	fmt.Println(t.GetNewsFeed(2))
}

type Twitter struct {
	g    *graph
	time int
}

func Constructor() Twitter {
	return Twitter{
		g: &graph{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.g.postNews(userId, tweetId, this.time)
	this.time++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	return this.g.getNews(userId)
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	this.g.addFollows(followerId, followeeId)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.g.unFollow(followerId, followeeId)
}

type graph struct {
	nodes []*node
}

type node struct {
	follows []*node
	posts   []post
	val     int
}

type post struct {
	id        int
	createdAt int
}

func (g *graph) getNode(val int) *node {
	for _, v := range g.nodes {
		if v.val == val {
			return v
		}
	}

	return nil
}

func (g *graph) addNode(val int) *node {
	n := &node{
		val: val,
	}

	g.nodes = append(g.nodes, n)

	return n
}

func (g *graph) addFollows(from, to int) {
	fromNode, toNode := g.getNode(from), g.getNode(to)
	if fromNode == nil {
		fromNode = g.addNode(from)
	}
	if toNode == nil {
		toNode = g.addNode(to)
	}
	for _, v := range fromNode.follows {
		if v.val == toNode.val {
			return
		}
	}
	fromNode.follows = append(fromNode.follows, toNode)
}

func (g *graph) unFollow(from, to int) {
	fromNode := g.getNode(from)

	for i := 0; i < len(fromNode.follows); i++ {
		if fromNode.follows[i].val == to {
			fromNode.follows = remove(fromNode.follows, i)
			break
		}
	}
}

func (g *graph) getNews(val int) []int {
	usr := g.getNode(val)
	if usr == nil {
		return nil
	}

	var posts []post
	var result []int

	for i := len(usr.follows) - 1; i >= 0; i-- {
		for j := len(usr.follows[i].posts) - 1; j >= 0; j-- {
			posts = append(posts, usr.follows[i].posts[j])
		}
	}

	for i := len(usr.posts) - 1; i >= 0; i-- {
		posts = append(posts, usr.posts[i])
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].createdAt > posts[j].createdAt
	})

	for i, v := range posts {
		if i == 10 {
			break
		}
		result = append(result, v.id)
	}

	return result
}

func (g *graph) postNews(val, id, createdAt int) {
	usr := g.getNode(val)
	if usr == nil {
		usr = g.addNode(val)
	}

	usr.posts = append(usr.posts, post{
		id:        id,
		createdAt: createdAt,
	})
}

func remove(slice []*node, s int) []*node {
	return append(slice[:s], slice[s+1:]...)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
