package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

	fmt.Println()

	a := [3]rune{'a', 'b', 'c'}
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%q\n", a)

	fmt.Println()
	m := map[string]int{"and": 1, "or": 2}
	fmt.Printf("%T\n", m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)

	fmt.Println()

	f := "a string"
	b := []byte(f)

	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%q\n", b)
}
