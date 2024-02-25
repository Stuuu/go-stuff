package main

import "fmt"

func main() {
	fmt.Println(doIt())
}

func doIt() *int {
	var b int

	return &b
}
