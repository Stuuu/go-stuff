package main

import "fmt"

func main() {
	var s []int
	var m map[string]int

	l := len(s)
	fmt.Println(l)

	i, ok := m["int"]
	fmt.Println(i, ok)

	for _, v := range s {
		fmt.Println(v)
	}

	// Make the zero value useful - Rob pike
}
