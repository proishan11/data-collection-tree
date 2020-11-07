package utils

import (
	"github.com/proishan11/data-collection-tree/models"
	"github.com/proishan11/data-collection-tree/tree"
)

func ResponseFromNode(node tree.Node) models.Res {
	var metrics = []struct {
		Key string `json:"key"`
		Val int32 `json:"val"`
	} {
		{
			"webreq",
			node.WebRequests,
		},

		{
			"timespent",
			node.TimeSpent,
		},
	}

	var dim = []struct {
		Key string `json:"key"`
		Val string `json:"val"`
	}{
		{
			"country",
			node.Country,
		},
	}

	return models.Res{
		Dim: dim,
		Metrics: metrics,
	}
}
