package domain

import "time"

type QueueRepository interface {
	CreateQueue(name string) error
	PutMessage(queueName string, message *Message) error
	GetMessage(queueName string, timeout time.Duration) (*Message, error)
}