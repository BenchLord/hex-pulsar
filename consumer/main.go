package main

import (
  "fmt"

  "github.com/BenchLord/hex-pulsar/consumer/ports"
  "github.com/BenchLord/hex-pulsar/consumer/adapters/pulsar"
)

type App struct {
  Queue ports.Queue
}

func NewApp(q ports.Queue) *App {
  return &App{q}
}

func main() {
  q, err := pulsar.NewQueue()
  if err != nil {
    panic(err)
  }

  app := NewApp(q)

  c, err := app.Queue.Subscribe(ports.SubscriptionOptions{
    Topic: "test-topic",
    Name: "test-consumer",
  })
  if err != nil {
    panic(err)
  }

  for {
    select {
      case msg := <-c:
        fmt.Println(string(msg.Payload))
        app.Queue.Ack(msg.MessageID)
      default:
    }
  }
}
