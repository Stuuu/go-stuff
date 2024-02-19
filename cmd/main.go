package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a string"

	x := len(s)

	fmt.Println(x)

	g := strings.Contains(s, "g")
	fmt.Println(g)
	l := strings.Contains(s, "x")
	fmt.Println(l)

	h := strings.HasPrefix(s, "a")
	fmt.Println(h)

	i := strings.Index(s, "string")
	fmt.Println(i)

	s = strings.ToUpper(s)
	fmt.Println(s)

}
