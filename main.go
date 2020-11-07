package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/proishan11/data-collection-tree/models"
	"github.com/proishan11/data-collection-tree/tree"
	"net/http"
)

var (
	t = tree.NewTree()
)

func insertHandler(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")

	var reqBody models.Request
	err := json.NewDecoder(req.Body).Decode(&reqBody)

	if err == nil {
		node := reqBody.Serialize()
		t.Insert(node)
		fmt.Println("Successfully inserted")
	} else {
		fmt.Println("Some error occurred")
	}
}

func queryHandler(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")

	var reqBody models.Query
	err := json.NewDecoder(req.Body).Decode(&reqBody)

	if err == nil {
		fmt.Println(reqBody.GetCountry())
		node := t.FindByCountry(reqBody.GetCountry())

		fmt.Println(node.TimeSpent)
		fmt.Println(node.WebRequests)
	}
}

func main() {
	const port string = ":8000"
	router := mux.NewRouter()
	router.HandleFunc("/v1/insert", insertHandler).Methods("POST")
	router.HandleFunc("/v1/query", queryHandler).Methods("GET")

	http.ListenAndServe(port, router)

	//t := tree.NewTree()
	//node := tree.NewNode()
	//
	//node.Country = "US"
	//node.Device = "Mobile"
	//node.WebRequests = 10
	//node.TimeSpent = 100
	//
	//n2 := tree.NewNode()
	//n2.Country = "US"
	//n2.WebRequests = 20
	//n2.TimeSpent = 90
	//n2.Device = "Laptop"
	//
	//n3 := tree.NewNode()
	//n3.Country = "US"
	//n3.WebRequests = 20
	//n3.TimeSpent = 90
	//n3.Device = "Laptop"
	//
	//t.Insert(*node)
	//t.Insert(*n2)
	//t.Insert(*n3)
	//
	//fmt.Println("Insertion Complete")
	//fmt.Println(t.FindByCountry("US").WebRequests)
	//fmt.Println(t.FindByCountry("US").TimeSpent)

}
