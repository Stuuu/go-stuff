package main

import "fmt"

func main() {
	fmt.Println(doIt())
}

func doIt() (a int) {
	defer func() {
		a = 2
	}()

	a = 1

	return
}
