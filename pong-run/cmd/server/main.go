package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/jjkoh95/onlyfap/pong-run/pkg/pongrun"
	"go.uber.org/zap"
)

var pr = pongrun.New("127.0.0.1", "6379", 10)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WSMsg struct {
	Action string
	Data   WSMsgData
}

type WSMsgData struct {
	Type   string `json:"type,omitempty"`
	Player int    `json:"player,omitempty"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allow", http.StatusMethodNotAllowed)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade websocket", http.StatusBadRequest)
		return
	}

	q := r.URL.Query()
	channelName := q.Get("id")

	// register
	pr.Subscribe(channelName, ws)
	count := len(pr.Subscriber.S[channelName].Conns)
	if count <= 2 {
		// player
		ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"player":%d}`, count)))
	}
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			zap.S().Error(err)
			return
		}
		pr.Publish(channelName, data)
	}

	// // write pump
	// go func() {}()
	// // read pump
	// go func() {}()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)

	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/ws", handleWS)
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
