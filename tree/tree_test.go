package tree

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestTreeFindByCountry(t *testing.T) {
	tree := NewTree()
	node := NewNode()

	node.Country = "US"
	node.Device = "Mobile"
	node.WebRequests = 10
	node.TimeSpent = 100

	tree.Insert(*node)

	res := tree.FindByCountry("US")

	if res.Country != node.Country && res.Device != node.Device && res.WebRequests != node.WebRequests &&
		res.TimeSpent != node.TimeSpent{
		t.Errorf("Insertion not verified")
	}

	root := tree.GetRoot()
	if root.TimeSpent != node.TimeSpent {
		t.Errorf("Root and Leave Timespent mismatch got %d, want %d", root.TimeSpent, node.TimeSpent)
	}

	if root.WebRequests != node.WebRequests {
		t.Errorf("Root and Leave Webrequests mismatch got %d, want %d", root.WebRequests, node.WebRequests)
	}
}

func TestTreeInsertConcurrently(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tree := NewTree()

	var countries = []string {"US", "IN", "HON", "CN", "UK", "RUS", "ESP", "BRA"}
	var devices = []string {"Apple", "MI", "NOKIA", "SAMSUNG", "HTC", "VIVO"}

	min := 30
	max := 150

	var usWebRequests int32 = 0
	var usTimeSpent int32 = 0

	var data []Node
	var totalWebRequest int32 = 0
	var totalTimeSpent int32 = 0

	for i:=0; i<100000; i++ {
		node := NewNode()
		node.WebRequests = int32(rand.Intn(max-min) + min)
		node.TimeSpent = int32(rand.Intn(max-min) + min)
		node.Device = devices[rand.Intn(len(devices)-1)]
		node.Country = countries[rand.Intn(len(countries)-1)]

		if node.Country == "US" {
			usWebRequests += node.WebRequests
			usTimeSpent += node.TimeSpent
		}

		totalWebRequest += node.WebRequests
		totalTimeSpent += node.TimeSpent
		data = append(data, *node)

	}
	var n sync.WaitGroup
	for _, d := range data {
		n.Add(1)
		go func(data Node) {
			tree.Insert(data)
			n.Done()
		}(d)
	}
	n.Wait()

	root := tree.GetRoot()

	if root.TimeSpent != totalTimeSpent {
		t.Errorf("Root Time Spent mismatch got %d, want %d", root.TimeSpent, totalTimeSpent)
	}

	if root.WebRequests != totalWebRequest {
		t.Errorf("Root web requests mismatch got %d, want %d", root.WebRequests, totalWebRequest)
	}

	usNode := tree.FindByCountry("US")
	if usNode.TimeSpent != usTimeSpent {
		t.Errorf("US Timespent mismatch got %d, want %d", usNode.TimeSpent, usTimeSpent)
	}

	if usNode.WebRequests != usWebRequests {
		t.Errorf("US webrequests mismatch got %d, want %d", usNode.WebRequests, usWebRequests)
	}
}

func TestTreeInsert(t *testing.T) {
	tree := NewTree()
	node := NewNode()

	node.Country = "US"
	node.Device = "Mobile"
	node.WebRequests = 10
	node.TimeSpent = 100

	n2 := NewNode()
	n2.Country = "US"
	n2.WebRequests = 20
	n2.TimeSpent = 90
	n2.Device = "Laptop"

	n3 := NewNode()
	n3.Country = "IN"
	n3.WebRequests = 20
	n3.TimeSpent = 90
	n3.Device = "Laptop"

	tree.Insert(*node)
	tree.Insert(*n2)
	tree.Insert(*n3)

	node = tree.FindByCountry("US")

	if node != nil {
		if node.TimeSpent != 190 {
			t.Errorf("got %d, want %d", node.TimeSpent, 190)
		}
		if node.WebRequests != 30 {
			t.Errorf("got %d, want %d", node.WebRequests, 30)
		}
	} else {
		t.Errorf("Node not found")
	}
}