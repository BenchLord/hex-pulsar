package main

import (
  "fmt"

  "github.com/BenchLord/hex-pulsar/consumer/ports"
  "github.com/BenchLord/hex-pulsar/consumer/adapter/pulsar"
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

  c, err := app.q.Subscribe(ports.SubscriptionOptions{
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
        q.Ack(msg.MessageID)
      default:
    }
  }
}
