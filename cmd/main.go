package main

import "fmt"

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	for _, item := range items {
		i := make([]byte, len(item)) // capture
		copy(i, item[:])             // copy prevents last mem reference from being set for all instances of item
		a = append(a, i)
	}

	fmt.Println(items)
	fmt.Println(a)
}
