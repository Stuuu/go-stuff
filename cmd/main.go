package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Page  int      `json:"page"`
	words []string `json:"words"`
}

func main() {

	r := &Response{Page: 1}
	j, _ := json.Marshal(r)
	fmt.Println(string(j))
	fmt.Printf("%#v\n", r)

	var r2 Response

	fmt.Printf("%#v\n", r2)
	_ = json.Unmarshal(j, &r2)

}
