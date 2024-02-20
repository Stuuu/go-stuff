package main

import "fmt"

func main() {
	p := map[string]int{} // non-nil but empty

	a := p["the"] // returns 0
	fmt.Println(a)

	b, ok := p["and"] // 0, false

	fmt.Println(b, ok)

	p["the"]++

	c, ok := p["the"] // 1, true

	fmt.Println(c, ok)

	if w, ok := p["the"]; ok {
		// we know w is not the default value
		fmt.Println(w)
	}

}
