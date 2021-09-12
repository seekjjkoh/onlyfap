package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	version = flag.String("version", "v1", "Version of app")
	port    = flag.String("port", "8081", "Expose port")
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf(`Version:%s`, *version)))
}

func main() {
	flag.Parse()
	http.HandleFunc("/", TestHandler)
	log.Println("Go server serving at port", *port)
	p := fmt.Sprintf(":%s", *port)
	http.ListenAndServe(p, nil)
}
