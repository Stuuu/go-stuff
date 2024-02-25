package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}
	b := a[:1]
	c := b[0:2:2]

	// all the same address
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("b[%p] = %[1]v\n", b)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c[0] = 9
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)

	c = append(c, 5)
	fmt.Printf("a[%p] = %v\n", &a, a)
	fmt.Printf("c[%p] = %[1]v\n", c)

}
