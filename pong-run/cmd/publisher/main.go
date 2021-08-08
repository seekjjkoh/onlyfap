package main

import (
	"fmt"

	"github.com/jjkoh95/onlyfap/pong-run/pkg/pongrun"
)

func main() {
	pr := pongrun.New("127.0.0.1", "6379", 10)
	// pr.Publish("channelName", []byte("hello world"))
	pr.SetState("test", []byte("aiushdiashd"))
	v, err := pr.GetState("adsadsad")
	fmt.Println(string(v), err)
	// fmt.Scanln()
}
