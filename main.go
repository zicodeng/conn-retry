package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/streadway/amqp"
)

const maxConnRetries = 5
const qName = "testq"

func connectToMQ(addr string) (*amqp.Connection, error) {
	mqURL := "amqp://" + addr
	var conn *amqp.Connection
	var err error
	for i := 1; i <= maxConnRetries; i++ {
		conn, err = amqp.Dial(mqURL)
		if err == nil {
			return conn, nil
		}
		log.Printf("error connecting to MQ server at %s: %s", mqURL, err)
		log.Printf("will attempt another connection in %d seconds", i*2)
		time.Sleep(time.Duration(i*2) * time.Second)
	}
	return nil, err
}

func listenToMQ(addr string) {
	conn, err := connectToMQ(addr)
	if err != nil {
		log.Fatalf("error connecting to MQ server: %s", err)
	}
	log.Printf("connected to MQ server")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("error opening channel: %v", err)
	}
	log.Println("created MQ channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(qName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("error declaring queue: %v", err)
	}
	log.Println("declared MQ queue")
	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("error listening to queue: %v", err)
	}
	log.Println("listening for new MQ messages...")
	for msg := range messages {
		log.Printf("new message id %s received from MQ", msg.MessageId)
	}
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	mqAddr := os.Getenv("MQADDR")
	if len(mqAddr) == 0 {
		log.Fatal("please set the MQADDR variable to the address of your MQ server")
	}

	go listenToMQ(mqAddr)

	log.Printf("web server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
