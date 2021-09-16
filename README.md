# Basit Http Request

[![GoDoc](https://godoc.org/github.com/ibnusurkati/timeformat?status.svg)](https://pkg.go.dev/github.com/ibnusurkati/timeformat?tab=doc)
[![License](https://img.shields.io/github/license/ibnusurkati/timeformat?style=plastic)](https://github.com/ibnusurkati/timeformat/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/ibnusurkati/timeformat.svg?style=flat-square)](https://github.com/ibnusurkati/timeformat/releases)


HTTP Client for golang, Inspired by [Javascript-axios](https://github.com/axios/axios). thanks to [Request](https://github.com/monaco-io/request) for being an inspiration. a small part of the code changes from [Request](https://github.com/monaco-io/request)

## Features
- Transform request and response data

## Installing

go mod:

```bash
go get github.com/ibnusurkati/basit
```

## Methods

- OPTIONS
- GET
- HEAD
- POST
- PUT
- DELETE
- TRACE
- CONNECT

## Example

```go
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
		DataType:     "json",
		ResponseType: "json",
	}

	resp := client.Exec(&result)
	if !resp.OK() {
		fmt.Println(resp.Error())
	}
	fmt.Println(result)
}
```