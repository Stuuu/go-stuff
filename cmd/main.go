package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a []int

	j1, _ := json.Marshal(a)
	fmt.Println(string(j1)) // null

	if j1 == nil {
		fmt.Println("j1 is nil")
	}

	b := []int{}

	j2, _ := json.Marshal(b)
	fmt.Println(string(j2)) // []

	if j2 == nil {
		fmt.Println("j2 is nil")
	}

}
