package main

import "fmt"

var w = [...]int{1, 2, 3}
var x = []int{0, 0, 0}

func main() {
	y := do(w, x)
	fmt.Println(w, x, y)

	z := w[:1]
	fmt.Println(z)
}

func do(a [3]int, b []int) []int {
	// a = b    // SYNTAX ERROR
	a[0] = 4 // w unchanged
	b[0] = 3 // x changed

	c := make([]int, 5) // []int{0,0,0,0,0}
	c[4] = 42
	copy(c, b) // copies only 3 elts

	return c
}
