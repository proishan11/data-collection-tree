package tree

import "sync"

type Node struct {

	mu sync.Mutex

	WebRequests int32
	TimeSpent int32
	Country string
	Device string
	Children map[string]*Node
}

func NewNode() *Node {
	return &Node{
		WebRequests: 0,
		TimeSpent: 0,
		Children: make(map[string]*Node),
		mu: sync.Mutex{},
	}
}