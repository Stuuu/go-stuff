package main

import "fmt"

func main() {

	n, err := fmt.Println("Hello, playground")
	fmt.Println(err)

	if _, err := fmt.Println(n); err != nil {
		fmt.Println(err)
	}
}
