package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

var (
	url   = flag.String("url", "http://localhost:8080", "URL")
	count = flag.Int("count", 1000, "Number of requests")
)

func main() {
	counter := make(map[string]int, 10)
	for i := 0; i < *count; i++ {
		res, _ := http.Get(*url)
		b, _ := io.ReadAll(res.Body)
		res.Body.Close()
		counter[string(b)]++
	}
	for k, v := range counter {
		fmt.Println(k, v)
	}
}
