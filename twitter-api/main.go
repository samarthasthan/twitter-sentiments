package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/samarthasthan/twitter-sentiments/types"
)

var sendInterval = time.Second

func main() {
	wsEndPoint := os.Getenv("WEBSOCKET_ENDPOINT")
	conn, _, err := websocket.DefaultDialer.Dial(wsEndPoint, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//Read Tweet CSV file
	file, err := os.Open("tweets.csv")
	if err != nil {
		log.Fatalln(err)
	}

	reader := csv.NewReader(file)

	tweets, err := reader.ReadAll()

	for _, data := range tweets {
		var tweet types.Tweet
		tweet.Username = data[4]
		tweet.Content = data[5]
		conn.WriteJSON(tweet)
		time.Sleep(sendInterval)
	}
}
