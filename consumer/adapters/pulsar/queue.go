package pulsar

import (
  "fmt"
  "strconv"

  "github.com/apache/pulsar-client-go/pulsar"
  "github.com/BenchLord/hex-pulsar/consumer/ports"
)

// Queue is a pulsar implementation of the Queue port.
type Queue struct {
  client pulsar.Client
  messages map[string]pulsar.ConsumerMessage
}

func NewQueue() (*Queue, error) {
  client, err := pulsar.NewClient(pulsar.ClientOptions{
    // This won't be hardcoded. NewQueue() will
    // need to take some config parameters.
    URL: "pulsar://localhost:6650",
  })
  if err != nil {
    return nil, fmt.Errorf("error creating pulsar client: %v", err)
  }

  q := &Queue{
    client: client,
    messages: make(map[string]pulsar.ConsumerMessage),
  }

  return q, nil
}

func (q *Queue) Subscribe(opts ports.SubscriptionOptions) (<-chan *ports.Message, error) {
  in := make(chan pulsar.ConsumerMessage)
  out := make(chan *ports.Message)

  consumer, err := q.client.Subscribe(pulsar.ConsumerOptions{
    Topic: opts.Topic,
    SubscriptionName: opts.Name,
    Type: pulsar.Shared,
    MessageChannel: in,
  })
  if err != nil {
    return nil, fmt.Errorf("error subscribing to pulsar topic: %v", err)
  }

  go func(q *Queue, in chan pulsar.ConsumerMessage, out chan *ports.Message) {
    count := 1
    defer consumer.Close()
    for {
      select {
        case msg := <-in:
          msgID := strconv.Itoa(count)
          // pass q in as param?
          q.messages[msgID] = msg
          out <- &ports.Message{
            MessageID: msgID,
            Payload: msg.Payload(),
          }

          count++
        default:
      }
    }
  }(q, in, out)

  return out, nil
}

func (q *Queue) Ack(messageID string) {
  m, ok := q.messages[messageID]
  if !ok {
    return
  }
  // will this work? a ConsumerMessage does
  // have a consumer object but it is the right
  // one?
  m.Ack(m)
  delete(q.messages, messageID)
}

func (q *Queue) Nack(messageID string) {
  m, ok := q.messages[messageID]
  if !ok {
    return
  }
  m.Nack(m)
  delete(q.messages, messageID)
}
