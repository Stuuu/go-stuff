package main

import "fmt"

func main() {
	var album1 = struct {
		title string
	}{
		"The White Album",
	}
	var album2 = struct {
		title string
	}{
		"The Black Album",
	}

	album1 = album2

	fmt.Println(album1, album2)
}
