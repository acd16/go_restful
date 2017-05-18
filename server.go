//REST api server for a simple string map

//TODO: better non existent url handling
//      better error handling
//      errors in json(?)

package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Pairs []Dict

type restErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

var data = make(map[string]string) //Global map to hold data

func GetDict(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var pairs Pairs //data list
	for k, v := range data {
		pairs = append(pairs, Dict{Key: k, Value: v})
	}
	if err := json.NewEncoder(w).Encode(pairs); err != nil { //add all keys
		panic(err)
	}
}

func GetDictKey(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	if val, ok := data[params["key"]]; ok { //check for key in dict and encode it
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Dict{Key: params["key"], Value: val}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(restErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func CreateDictKey(w http.ResponseWriter, req *http.Request) {
	var dict Dict

	if err := json.NewDecoder(req.Body).Decode(&dict); err != nil {
		panic(err) //Error if input format is wrong
	}
	data[dict.Key] = dict.Value //add key from input
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(Dict{Key: dict.Key, Value: dict.Value}); err != nil {
		panic(err)
	}
}

func DeleteDictKey(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	delete(data, params["key"])
}

func UpdateDictKey(w http.ResponseWriter, req *http.Request) {
	var dict Dict

	if err := json.NewDecoder(req.Body).Decode(&dict); err != nil {
		panic(err) //Error if input format is wrong
		return
	}
	if _, ok := data[dict.Key]; ok { //check for key in dict and update it
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		data[dict.Key] = dict.Value
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(restErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
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
