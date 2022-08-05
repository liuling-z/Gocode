package main

import (
	real2 "Gocode/retriever/real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = real2.Retriever{}
	fmt.Println(download(r))
}
