//client for simple REST server

//TODO:add flags
//	   report errors from server
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Get a specific key based on provided key
func getKey(key string) {
	reqStr := "http://localhost:8080/v1/dict/" + key
	req, err := http.NewRequest("GET", reqStr, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	var out Dict
	json.NewDecoder(resp.Body).Decode(&out)

	res, _ := json.Marshal(out)
	fmt.Println(string(res))
}

//add a key, value pair
func createKey(key, value string) {
	data := Dict{Key: key, Value: value}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(dataBytes)

	req, err := http.NewRequest("POST", "http://localhost:8080/v1/dict/add/", body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return
}

//List all keys
func getAllKeys() {
	resp, err := http.Get("http://localhost:8080/v1/dict/")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		panic(err)
	}
}

//Delete specific key
func deleteKey(key string) {
	reqStr := "http://localhost:8080/v1/dict/delete/" + key
	req, err := http.NewRequest("DELETE", reqStr, nil)

	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
}

func updateKey(key, value string) {
	reqStr := "{\"key\":\"" + key + "\", \"value\":\"" + value + "\"}"
	body := strings.NewReader(reqStr)
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/dict/update/", body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
}

func main() {
	//add, update, delete, get, getall
	addPtr := flag.String("add", "", "key,value to be added")
	updatePtr := flag.String("update", "", "key,value to be updated")
	deletePtr := flag.String("delete", "", "key to be deleted")
	getPtr := flag.String("get", "", "key to retrieve")
	getAllPtr := flag.Bool("getAll", false, "get all keys")

	flag.Parse()

	if *addPtr != "" {
		s := strings.Split(*addPtr, ",")
		createKey(s[0], s[1])
	}

	if *updatePtr != "" {
		s := strings.Split(*updatePtr, ",")
		updateKey(s[0], s[1])
	}

	if *deletePtr != "" {
		deleteKey(*deletePtr)
	}

	if *getPtr != "" {
		getKey(*getPtr)
	}

	if *getAllPtr == true {
		getAllKeys()
	}
}
