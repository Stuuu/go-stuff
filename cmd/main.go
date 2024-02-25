package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3}
	b := a[:1]

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := b[0:2]

	fmt.Println("c = ", c)

	fmt.Println(len(b))
	fmt.Println(cap(b))

	fmt.Println(len(c))
	fmt.Println(cap(c))

	d := c[0:1:1] // [i:j:k] len j-i cap k -i

	fmt.Println("d = ", d)

	fmt.Println(len(d))
	fmt.Println(cap(d))
}
