package main

import "fmt"

func main() {

	var i int
	for i = 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}

		do(v)
	}
}

func do(d func()) {
	d()
}
