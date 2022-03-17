package ports

// Queue defines behavior needed for a message queue.
type Queue interface {
  Subscribe(SubscriptionOptions) (<-chan *Message, error)
  Ack(string)
  Nack(string)
}

// Message defines the data that will be set through the channel created
// when subscribing to a producer.
type Message struct {
  string MessageID
  []byte Payload
}

// SubscriptionOptions are used to subscribe to a producer.
type SubscriptionOptions struct {
  string Topic
  string Name
}
