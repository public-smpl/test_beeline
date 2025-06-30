package domain

import "time"

type Message struct {
	Data      string    `json:"message"`
	Timestamp time.Time
}

type QueueConsumer struct {
	ResponseChan chan *Message
	Timeout      time.Duration
	RequestTime  time.Time
}

type Queue struct {
	Name      string
	Messages  []*Message
	Consumers []*QueueConsumer
}

func NewQueue(name string) *Queue {
	return &Queue{
		Name:      name,
		Messages:  make([]*Message, 0),
		Consumers: make([]*QueueConsumer, 0),
	}
}

func NewMessage(data string) *Message {
	return &Message{
		Data:      data,
		Timestamp: time.Now(),
	}
}

func NewQueueConsumer(timeout time.Duration) *QueueConsumer {
	return &QueueConsumer{
		ResponseChan: make(chan *Message, 1),
		Timeout:      timeout,
		RequestTime:  time.Now(),
	}
}