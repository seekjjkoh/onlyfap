package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

var (
	url         = flag.String("url", "http://localhost:8080", "URL")
	count       = flag.Int("count", 1000, "Number of requests")
	headerName  = flag.String("headerName", "ab-variation", "Custom header name")
	headerValue = flag.String("headerValue", "service1private", "Custom header value")
)

func main() {
	counter := make(map[string]int, 10)
	httpClient := http.Client{}

	for i := 0; i < *count; i++ {
		req, _ := http.NewRequest("GET", *url, nil)
		req.Header.Add(*headerName, *headerValue)
		res, _ := httpClient.Do(req)
		b, _ := io.ReadAll(res.Body)
		res.Body.Close()
		counter[string(b)]++
	}
	for k, v := range counter {
		fmt.Println(k, v)
	}
}
