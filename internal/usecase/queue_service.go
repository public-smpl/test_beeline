package usecase

import (
	"time"

	"queue-broker/internal/domain"
)

type QueueServiceImpl struct {
	repo domain.QueueRepository
}

func NewQueueService(repo domain.QueueRepository) domain.QueueService {
	return &QueueServiceImpl{
		repo: repo,
	}
}

func (s *QueueServiceImpl) PutMessage(queueName string, messageData string) error {
	message := domain.NewMessage(messageData)
	return s.repo.PutMessage(queueName, message)
}

func (s *QueueServiceImpl) GetMessage(queueName string, timeout time.Duration) (*domain.Message, error) {
	return s.repo.GetMessage(queueName, timeout)
}