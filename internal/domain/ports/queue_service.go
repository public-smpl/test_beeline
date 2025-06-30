package domain

import "time"

type QueueService interface {
	PutMessage(queueName string, messageData string) error
	GetMessage(queueName string, timeout time.Duration) (*Message, error)
}