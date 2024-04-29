package main

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

	comic_map := make(map[int]Comic)
	// xkcd is surprisingly not 0 indexed
	comic_count := 1
	for {

		comic, err := fetchComicJson(comic_count)
		if err != nil {
			fmt.Println(err)
			break;
		}
		comic_map[comic_count] = comic
		comic_count++
		if comic_count == 404 {
			comic_count++	
		}
	}	
	
	// Create a file for IO
	encodeFile, err := os.Create("comics.gob")
	if err != nil {
		panic(err)
	}

	// Since this is a binary format large parts of it will be unreadable
	encoder := gob.NewEncoder(encodeFile)

	// Write to the file
	if err := encoder.Encode(comic_map); err != nil {
		panic(err)
	}
	fmt.Println("Files fetched")
	encodeFile.Close()

}

func fetchComicJson(comic_num int) (Comic, error) {

	request_url := "https://xkcd.com/" + strconv.Itoa(comic_num) + "/info.0.json"
	fmt.Println("fetched", request_url)

	req, err := http.NewRequest(http.MethodGet, request_url, nil)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusNotFound {
		return Comic{}, errors.New("Comic Not Found: " + strconv.Itoa(comic_num))
	}
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	//We Read the response body on the line below.
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var comic1 Comic
	err = json.Unmarshal(body, &comic1)
	if err != nil {
		log.Fatalln(err)
	}
	return comic1, nil
}
