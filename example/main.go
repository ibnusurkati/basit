package main

import (
	"fmt"

	"github.com/ibnusurkati/basit"
)

type Coba struct {
	Origin string `json:"origin"`
	Url    string `json:"url"`
}

func main() {
	var result interface{}
	client := basit.Instance{
		Url:    "http://httpbin.org/post",
		Method: "POST",
		Data: Coba{
			Origin: "bismillah",
			Url:    "Alhamdulillah",
		},
		Headers: map[string]string{
			"public-key": "hello cuy",
			"User-Agent": "irx",
		},
		Query: map[string]string{
			"public-key": "hello cuy",
			"User-Agent": "irx",
		},
		DataType:     "json",
		ResponseType: "json",
	}

	resp := client.Exec(&result)
	if !resp.OK() {
		fmt.Println(resp.Error())
	}
	fmt.Println(result)
}
