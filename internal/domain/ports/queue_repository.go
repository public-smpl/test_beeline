package ports

import (
	"time"
	"queue-broker/internal/domain"
)

type QueueRepository interface {
	CreateQueue(name string) error
	PutMessage(queueName string, message *domain.Message) error
	GetMessage(queueName string, timeout time.Duration) (*domain.Message, error)
}