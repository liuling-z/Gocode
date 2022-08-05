package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//Go语言数据I/O对象及操作
func main() {
	data := []byte("Hello GoLand")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:]), n, err)
}
