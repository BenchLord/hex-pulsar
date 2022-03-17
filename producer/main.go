package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		log.Fatal("error creating pulsar client: %v", err)
	}

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "test-topic",
	})
	if err != nil {
		log.Fatal("error creating producer: %v", err)
	}
	defer producer.Close()

	count := 1
	for {
		if _, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(fmt.Sprintf("message number %d", count)),
		}); err != nil {
			log.Fatal("error sending producer message: %v", err)
		}
		count++

		time.Sleep(2 * time.Second)
	}
}
