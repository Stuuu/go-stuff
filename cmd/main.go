package main

import "fmt"

func main() {
	a, b := 12, 345
	c, d := 1.2, 3.45

	fmt.Println()

	fmt.Printf("|%9d| |%9d|", a, b)
	fmt.Println()
	fmt.Printf("|%09d| |%09d|", a, b)
	fmt.Println()
	fmt.Printf("|%-9d| |%-9d|", a, b)

	fmt.Println()

	fmt.Printf("|%9f| |%9.2f|", c, d)
}
