package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)



func main(){
	
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.ListenAndServe(":8080", nil)
}
var items = map[string]int {
	"corndog": 2,
	"hotdog": 1,
}
func createHandler(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	
	
	if item == "" {
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("No item provided on create"))
		return
	}
	
	if items[item] != 0 {
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("Item already exist can not create: " + item))
		return
	}
	if price == "" {
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("No price provided on create"))
		return
	}
	
	items[item], _ = strconv.Atoi(price)	
	item_json, _ := json.Marshal(items)
	fmt.Fprintf(w, string(item_json))
}
func readHandler(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	
	if item == "" {
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("No item provided for read"))
		return
	}
	item_json, _ := json.Marshal(items[item])
	fmt.Fprintf(w, string(item_json))
}
func updateHandler(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	
	if item == "" {
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("No item provided for update"))
		return
	}
	if price == "" {
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("No price provided for update: "))
		return
	}
	
	if items[item] == 0 {
		w.WriteHeader(http.StatusInternalServerError)
    	w.Write([]byte("No item found for update: "))
		return
	}
	
	items[item], _ = strconv.Atoi(price)
	item_json, _ := json.Marshal(items)
	fmt.Fprintf(w, string(item_json))
}
func deleteHandler(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	delete(items, item)
	item_json, _ := json.Marshal(items)
	fmt.Fprintf(w, string(item_json))
}