package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/samarthasthan/twitter-sentiments/types"
)

func main() {
	port := os.Getenv("WEBSOCKET_PORT")
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/ws", recv.handleWebSocket)
	http.ListenAndServe(port, nil)
}

type Datareceiver struct {
	conn *websocket.Conn
}

func NewDataReceiver() (*Datareceiver, error) {
	return &Datareceiver{}, nil
}

func (dr *Datareceiver) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	var upgrade = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}

	dr.conn = conn

	go dr.wsReceiverLoop()
}

func (dr *Datareceiver) wsReceiverLoop() {
	for {
		var tweet types.Tweet
		dr.conn.ReadJSON(&tweet)
		fmt.Printf("%+v\n", tweet)
	}
}
