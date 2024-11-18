package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id  int64
	Msg string
}

func main() {
	var i int64 = 0
	c1 := make(chan Message)
	c2 := make(chan Message)

	//RabbitMQ
	go func() {

		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}

	}()

	//Kafka
	go func() {

		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}

	}()

	// for i := 0; i < 2; i++ {
	for {

		select {
		case msg := <-c1:
			fmt.Printf("Received from RabbitMQ: ID %d - %s\n", msg.id, msg.Msg)

		case msg := <-c2:
			fmt.Printf("Received from Kafka: ID %d - %s\n", msg.id, msg.Msg)
		case <-time.After(time.Second * 3):
			println("Timeout")
			// default:
			// 	println("default")

		}
	}

	// }
}
