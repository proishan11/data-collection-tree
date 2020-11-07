package models

import "github.com/proishan11/data-collection-tree/tree"

type Request struct {
	Dim []struct {
		Key string `json:"key"`
		Val string `json:"val"`
	} `json:"dim"`
	Metrics []struct {
		Key string `json:"key"`
		Val int32    `json:"val"`
	} `json:"metrics"`
}

func (r *Request) Serialize() tree.Node {

	n := tree.NewNode()

	for _, d := range r.Dim {
		if d.Key == "device" {
			n.Device = d.Val
		}
		if d.Key == "country" {
			n.Country = d.Val
		}
	}

	for _, m := range r.Metrics {
		if m.Key == "webreq" {
			n.WebRequests = m.Val
		}

		if m.Key == "timespent" {
			n.TimeSpent = m.Val
		}
	}
	return *n
}