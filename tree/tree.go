package tree

import "fmt"

type tree struct {
	rootNode *Node
}

type Tree interface {
	Insert(newNode Node)
	FindByCountry(country string) *Node
}

func(t *tree) Insert(newNode Node) {

	t.rootNode.mu.Lock()
	t.rootNode.WebRequests += newNode.WebRequests
	t.rootNode.TimeSpent += newNode.TimeSpent
	t.rootNode.mu.Unlock()

	c, ok := t.rootNode.Children[newNode.Country]
	if !ok {
		fmt.Println("Country Node not present")
		cNode := NewNode()
		cNode.TimeSpent = newNode.TimeSpent
		cNode.WebRequests = newNode.WebRequests

		dNode := NewNode()
		dNode.TimeSpent = newNode.TimeSpent
		dNode.WebRequests = newNode.WebRequests

		t.rootNode.mu.Lock()
		t.rootNode.Children[newNode.Country] = cNode
		t.rootNode.Children[newNode.Country].Children[newNode.Device] = dNode
		t.rootNode.mu.Unlock()
	} else {

		c.mu.Lock()
		c.WebRequests += newNode.WebRequests
		c.TimeSpent += newNode.TimeSpent
		c.mu.Unlock()

		d, ok := t.rootNode.Children[newNode.Country].Children[newNode.Device]

		if ok {
			fmt.Println("Device Node present")
			d.mu.Lock()
			d.TimeSpent += newNode.TimeSpent
			d.WebRequests += newNode.WebRequests
			d.mu.Unlock()
		} else {
			fmt.Println("Device Node not present")

			dNode := NewNode()

			dNode.TimeSpent = newNode.TimeSpent
			dNode.WebRequests = newNode.WebRequests
			dNode.Device = newNode.Device

			t.rootNode.Children[newNode.Country].mu.Lock()
			t.rootNode.Children[newNode.Country].Children[newNode.Device] = dNode
			t.rootNode.Children[newNode.Country].mu.Unlock()
		}
	}
}

func(t *tree) FindByCountry(country string) *Node {
	return t.rootNode.Children[country]
}


func NewTree() Tree {
	newNode := NewNode()

	return &tree {
		rootNode: newNode,
	}
}