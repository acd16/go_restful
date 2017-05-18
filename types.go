//Common types used by server and client

package main 

type Dict struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type restErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type Pairs []Dict
