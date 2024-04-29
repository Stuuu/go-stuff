package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// declaring a struct
type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	Safe_Title string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func main() {

	
	search_terms := os.Args[1:]
	if len(search_terms) == 0  {
		fmt.Println("No search terms provided, at least 1 required")
		os.Exit(0)
	}
	
	
	comics_maps := decodeComicMap()
	
	search_results := make(map[int]Comic)
	for  _, comic := range comics_maps {
		for _, term := range search_terms {
			if strings.Contains(comic.Title, term) || strings.Contains(comic.Transcript, term) {
				search_results[comic.Num] = comic
			}
		}
	}
	
	for _, result := range search_results {
		fmt.Println(result.Title)
		fmt.Println(result.Day + "/" + result.Month + "/" + result.Year)
		fmt.Println("https://xkcd.com/" + strconv.Itoa(result.Num))
		fmt.Println()
	}

	fmt.Println("Results Found: ", len(search_results))
}


func decodeComicMap() map[int]Comic {
	// Open a RO file
	decodeFile, err := os.Open("comics.gob")
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()

	// Create a decoder
	decoder := gob.NewDecoder(decodeFile)

	// Place to decode into
	comics := make(map[int]Comic)

	// Decode -- We need to pass a pointer otherwise accounts2 isn't modified
	decoder.Decode(&comics)
	
	return comics
}