package main

import "fmt"

func main() {

	things := []int{1, 2, 3, 4, 5, 6, 7}

	for _, thing := range things {
		thing++
	}

	fmt.Println(things)

	for i := range things {
		things[i]++
	}

	fmt.Println(things)
}
