package ports

import (
	"time"
	"queue-broker/internal/domain"
)

type QueueService interface {
	PutMessage(queueName string, messageData string) error
	GetMessage(queueName string, timeout time.Duration) (*domain.Message, error)
}