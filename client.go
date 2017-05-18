//client for simple REST server

//TODO:add flags 
//	   report errors from server
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"os"
	"strings"
)

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

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
	reqStr := "{\"key\":"+key+", \"value\":"+value+"}"
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
	getKey("foo")
	createKey("abc", "def")
	getKey("abc")
	getAllKeys()
	deleteKey("foo")
	getAllKeys()
	updateKey("abc", "deg")
	getAllKeys()
}
