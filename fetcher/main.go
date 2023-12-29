package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/samarthasthan/twitter-sentiments/producer"
	"github.com/samarthasthan/twitter-sentiments/types"
)

func main() {
	port := os.Getenv("WEBSOCKET_PORT")
	p := producer.NewKafkaProducer()

	recv, err := NewDataReceiver(p)
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/ws", recv.handleWebSocket)
	http.ListenAndServe(port, nil)
}

type Datareceiver struct {
	conn     *websocket.Conn
	producer producer.DataProducer
}

func NewDataReceiver(p producer.DataProducer) (*Datareceiver, error) {
	return &Datareceiver{producer: p}, nil
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
		err := dr.conn.ReadJSON(&tweet)
		if err != nil {
			log.Fatalln(err)
		} else {
			dr.producer.Produce(&tweet)
		}
	}
}
