package main

import (
	"log"
	"net/http"
)

const PORT = ":1800"

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	http.HandleFunc("/", TestHandler)
	log.Println("Go server serving at port", PORT)
	http.ListenAndServe(PORT, nil)
}
