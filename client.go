//client for simple REST server

//TODO: reject invalid inputs

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

const serverUrl = "http://localhost:8080"

//Get a specific key based on provided key
func getKey(key string) {
	reqStr := serverUrl + "/v1/dict/" + key
	req, err := http.NewRequest("GET", reqStr, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	fmt.Println("Status: " + resp.Status)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		var out Dict
		json.NewDecoder(resp.Body).Decode(&out)

		res, _ := json.Marshal(out)
		fmt.Println(string(res))
	} else {
		var out restErr
		json.NewDecoder(resp.Body).Decode(&out)

		res, _ := json.Marshal(out)
		fmt.Println(string(res))
	}
}

//add a key, value pair
func createKey(key, value string) {
	data := Dict{Key: key, Value: value}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(dataBytes)

	reqUrl := serverUrl + "/v1/dict/add/"
	req, err := http.NewRequest("POST", reqUrl, body)
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
	fmt.Println("Status: " + resp.Status)

	return
}

//List all keys
func getAllKeys() {
	reqUrl := serverUrl + "/v1/dict/"
	resp, err := http.Get(reqUrl)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println("Status: " + resp.Status)
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		panic(err)
	}
}

//Delete specific key
func deleteKey(key string) {
	reqUrl := serverUrl + "/v1/dict/delete/" + key
	req, err := http.NewRequest("DELETE", reqUrl, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Status: " + resp.Status)
}

func updateKey(key, value string) {
	reqStr := "{\"key\":\"" + key + "\", \"value\":\"" + value + "\"}"
	body := strings.NewReader(reqStr)
	reqUrl := serverUrl + "/v1/dict/update/"
	req, err := http.NewRequest("PUT", reqUrl, body)
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
	fmt.Println("Status: " + resp.Status)
	if resp.StatusCode != 200 {
		var out restErr
		json.NewDecoder(resp.Body).Decode(&out)

		res, _ := json.Marshal(out)
		fmt.Println(string(res))
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
