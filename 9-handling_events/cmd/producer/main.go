package main

import "github.com/BielPinto/curso_go/9-handling_events/pkg/events/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)

	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello world", "amq.direct")
}
