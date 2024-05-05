package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	l := len(p)
	*b += ByteCounter(l)
	return l, nil
}

func main() {
	var c ByteCounter
	
	f1, _ := os.Open("a.txt")
	f2 := &c 
	
	n, _ := io.Copy(f2, f1)
	
	fmt.Println("copied", n, "bytes")
	fmt.Println(c)
}