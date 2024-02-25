package main

import (
	"fmt"
)

func main() {

	u := make([]int, 5)

	b := make([]int, 0, 5)

	fmt.Println(u)

	u = append(u, 5)
	b = append(b, 5)

	fmt.Println(u)
	fmt.Println(b)

}
