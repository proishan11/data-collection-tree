package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/proishan11/data-collection-tree/models"
	"github.com/proishan11/data-collection-tree/tree"
	"github.com/proishan11/data-collection-tree/utils"
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
		response.WriteHeader(http.StatusOK)
	} else {
		response.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Some error occurred ", err)
	}
}

func queryHandler(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-type", "application/json")

	var reqBody models.Query
	err := json.NewDecoder(req.Body).Decode(&reqBody)

	if err == nil {
		fmt.Println(reqBody.GetCountry())
		node := t.FindByCountry(reqBody.GetCountry())

		if node != nil {
			res := utils.ResponseFromNode(*node)
			response.WriteHeader(http.StatusOK)
			json.NewEncoder(response).Encode(res)
		} else {
			response.WriteHeader(http.StatusOK)
		}
	} else {
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	const port string = ":8000"
	router := mux.NewRouter()
	router.HandleFunc("/v1/insert", insertHandler).Methods("POST")
	router.HandleFunc("/v1/query", queryHandler).Methods("GET")

	http.ListenAndServe(port, router)
}
