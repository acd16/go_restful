//REST api server for a simple string map

//TODO: better non existent url handling
//      better error handling
//      errors in json(?)

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var data = make(map[string]string) //Global map to hold data

func GetDict(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data) //add all keys
}

func GetDictKey(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if val, ok := data[params["key"]]; ok { //check for key in dict and encode it
		//TODO: simple dict
		json.NewEncoder(w).Encode(Dict{Key: params["key"], Value: val})
		return
	}
	fmt.Fprintf(w, "Error, key not found\n")
}

func CreateDictKey(w http.ResponseWriter, req *http.Request) {
	var dict Dict

	if err := json.NewDecoder(req.Body).Decode(&dict); err != nil {
		fmt.Fprintf(w, "Error\n") //Error if input format is wrong
		return
	}
	data[dict.Key] = dict.Value //add key from input
}

func DeleteDictKey(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	delete(data, params["key"])
}

func UpdateDictKey(w http.ResponseWriter, req *http.Request) {
	var dict Dict

	if err := json.NewDecoder(req.Body).Decode(&dict); err != nil {
		fmt.Fprintf(w, "Error\n") //Error if input format is wrong
		return
	}
	if _, ok := data[dict.Key]; ok { //check for key in dict and update it
		data[dict.Key] = dict.Value
		return
	}
	fmt.Fprintf(w, "Error, key not found\n")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/dict/", GetDict).Methods("GET")
	router.HandleFunc("/v1/dict/{key}", GetDictKey).Methods("GET")
	router.HandleFunc("/v1/dict/add/", CreateDictKey).Methods("POST")
	router.HandleFunc("/v1/dict/update/", UpdateDictKey).Methods("PUT")
	router.HandleFunc("/v1/dict/delete/{key}", DeleteDictKey).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
