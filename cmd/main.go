package main

import "fmt"

func main() {
	a := "a"
	b := "a"
	if a == b {
		fmt.Println("equals")
	} else {
		fmt.Println("does not equal")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("(%d, %d)\n", i, i*i)
	}

}
