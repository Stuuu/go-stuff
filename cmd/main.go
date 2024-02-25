package main

import (
	"fmt"
	"strconv"
)

func printLater(s string) {

	fmt.Println(s)
}

func main() {
	s := "this is getting printed later"
	defer printLater(s)

	for i := 0; i < 100; i++ {
		s = s + strconv.Itoa(i)
		defer printLater(s)
	}

}
