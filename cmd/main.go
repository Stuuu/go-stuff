package main

import (
	"fmt"
	"math/rand"
)

func getNum() int {
	return rand.Intn(10-0) + 0
}
func main() {
	switch a := getNum(); a {
	case 0, 1, 2:
		fmt.Println("case 1", a)
	case 3, 4, 5, 6, 7, 8:
		fmt.Println("case 2", a)
	default:
		fmt.Println("case 3", a)
	}

	a := getNum()
	// Switch on true
	switch {
	case a <= 2:
		fmt.Println("less than 2", a)
	case a <= 8:
		fmt.Println("greater than 2 less than 8", a)
	}

}
