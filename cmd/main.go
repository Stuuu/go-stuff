package main

import "fmt"

func main() {
	var m = map[string]int{
		"and": 1,
		"the": 1,
		"or":  2,
	}

	var n map[string]int

	// b := m == n   // SYntax Error
	c := n == nil // true
	d := len(m)   // 3
	// e := cap(m)   // Type mismatch you can't ask for the capacity of a map

	fmt.Println(c, d)

}
