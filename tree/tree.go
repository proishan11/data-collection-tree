package tree

import "fmt"

type tree struct {
	rootNode Node
}

type Tree interface {
	Insert(newNode Node)
}

func(t *tree) Insert(newNode Node) {

	// acquire root Node lock here
	t.rootNode.mu.Lock()
	t.rootNode.WebRequests += newNode.WebRequests
	t.rootNode.TimeSpent += newNode.TimeSpent
	t.rootNode.mu.Unlock()
	// free rootNode lock here

	c, ok := t.rootNode.Children[newNode.Country]
	if !ok {

		fmt.Println("Country Node not present")

		// acquire Node lock
		t.rootNode.mu.Lock()
		t.rootNode.Children[newNode.Country] = newNode
		t.rootNode.Children[newNode.Country].Children[newNode.Device] = newNode
		t.rootNode.mu.Unlock()
		// free Node lock

	} else {

		// acquire Lock

		c.WebRequests += newNode.WebRequests
		c.TimeSpent += newNode.TimeSpent
		t.rootNode.Children[newNode.Country] = c

		d, ok := t.rootNode.Children[newNode.Country].Children[newNode.Device]

		if ok {
			fmt.Println("Device Node present")
			d.TimeSpent += newNode.TimeSpent
			d.WebRequests += newNode.WebRequests

			// acquire lock
			//t.rootNode.Children[newNode.Country].Children[newNode.Device].mu.Lock()

			// There is no need to lock this I guess because of the initial nil value
			// Should I lock the parent pointer here?
			t.rootNode.Children[newNode.Country].Children[newNode.Device] = d
			//t.rootNode.Children[newNode.Country].Children[newNode.Device].mu.Unlock()
			// free lock

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

func NewTree() Tree {
	newNode := NewNode()

	return &tree {
		rootNode: newNode,
	}
}