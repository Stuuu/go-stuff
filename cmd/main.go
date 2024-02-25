package main

import "fmt"

func main() {
	f := fib()

	for x := f(); x < 100; x = f() {
		fmt.Println(x)
	}

}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		fmt.Println(a, b)
		a, b = b, a+b
		return b
	}
}
