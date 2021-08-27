package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/jjkoh95/onlyfap/pong-run/pkg/pongrun"
	"go.uber.org/zap"
)

var pr = pongrun.New(os.Getenv("REDISHOST"), os.Getenv("REDISPORT"), 10)

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
	countB, err := pr.GetState(channelName)
	var count byte
	if err != nil {
		zap.S().Error(err)
		count = 1
		pr.SetState(channelName, []byte{1})
	} else {
		count = countB[0] + 1
		if count == 2 {
			pr.SetState(channelName, []byte{count})
		}
	}
	zap.S().Info("count", count)
	ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"player":%d}`, count)))
	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			zap.S().Error(err)
			if websocket.IsCloseError(err, websocket.CloseGoingAway) || err == io.EOF {
				pr.DeRegisterWs(channelName, ws)
				return
			}
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
