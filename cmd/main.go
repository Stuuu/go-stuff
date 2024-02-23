package main

import "fmt"

func main() {
	a, b := 12, 345
	c, d := 1.2, 3.45

	fmt.Printf("%d %d\n", a, b)
	fmt.Printf("%X %X\n", a, b)
	fmt.Printf("%#x %#x\n", a, b)
	fmt.Printf("%f %.5f\n", c, d)

}
