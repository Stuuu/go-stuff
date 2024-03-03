package main

import "fmt"

func main() {
	v1 := struct {
		X int `json:"foo"`
	}{1}

	v2 := struct {
		X int `json:"foo"`
	}{2}

	v1 = v2

	fmt.Println(v1)
}
