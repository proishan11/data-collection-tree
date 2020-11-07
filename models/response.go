package models

type Res struct {
	Dim []struct {
		Key string `json:"key"`
		Val string `json:"val"`
	} `json:"dim"`
	Metrics []struct {
		Key string `json:"key"`
		Val int32    `json:"val"`
	} `json:"metrics"`
}