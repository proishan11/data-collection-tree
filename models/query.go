package models

type Query struct {
	Dim []struct {
		Key string `json:"key"`
		Val string `json:"val"`
	} `json:"dim"`
}

func (q *Query) GetCountry() string {

	for _, c := range q.Dim {
		if c.Key == "country" {
			return c.Val
		}
	}
	return ""
}