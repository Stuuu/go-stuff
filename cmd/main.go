package main

import "fmt"

func main() {
	var m map[string]int
	p := make(map[string]int)

	a := p["the"]
	b := m["the"]
	fmt.Println(a, b)
	p["and"] = 1 // PANIC - nil map
	m = p
	m["and"]++
	c := p["and"]

	fmt.Println(m, p, c)
}
