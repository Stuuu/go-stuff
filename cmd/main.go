package main

import "fmt"

func main() {
	s := "élite"

	fmt.Printf("%8T %[1]v %d\n", s, len(s))
	fmt.Printf("%8T %[1]v\n", []rune(s))
	b := []byte(s)
	fmt.Printf("%8T %[1]v %d\n", b, len(b))

	s2 := "the quick brown fox"

	a := len(s2)
	fmt.Println(a)
	b2 := s2[:3]
	fmt.Println(b2)
	c := s2[4:9]
	fmt.Println(c)
	d := s2[:4] + "slow" + s2[9:]
	fmt.Println(d)

	s2 += "es"

	fmt.Println(s2)

}
