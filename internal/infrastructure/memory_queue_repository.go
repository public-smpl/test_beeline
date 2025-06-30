package infrastructure

import (
	"errors"
	"sync"
	"time"

	"queue-broker/internal/domain"
)

type MemoryQueueRepository struct {
	queues       map[string]*domain.Queue
	mutex        sync.RWMutex
	maxQueues    int
	maxMessages  int
}

func NewMemoryQueueRepository(maxQueues, maxMessages int) *MemoryQueueRepository {
	return &MemoryQueueRepository{
		queues:      make(map[string]*domain.Queue),
		maxQueues:   maxQueues,
		maxMessages: maxMessages,
	}
}

func (r *MemoryQueueRepository) CreateQueue(name string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.queues[name]; exists {
		return nil
	}

	if len(r.queues) >= r.maxQueues {
		return errors.New("maximum number of queues reached")
	}

	r.queues[name] = domain.NewQueue(name)
	return nil
}

func (r *MemoryQueueRepository) PutMessage(queueName string, message *domain.Message) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	queue, exists := r.queues[queueName]
	if !exists {
		if len(r.queues) >= r.maxQueues {
			return errors.New("maximum number of queues reached")
		}
		queue = domain.NewQueue(queueName)
		r.queues[queueName] = queue
	}

	if len(queue.Messages) >= r.maxMessages {
		return errors.New("queue is full")
	}

	if len(queue.Consumers) > 0 {
		consumer := queue.Consumers[0]
		queue.Consumers = queue.Consumers[1:]
		consumer.ResponseChan <- message
		return nil
	}

	queue.Messages = append(queue.Messages, message)
	return nil
}

func (r *MemoryQueueRepository) GetMessage(queueName string, timeout time.Duration) (*domain.Message, error) {
	r.mutex.Lock()
	queue, exists := r.queues[queueName]
	if !exists {
		queue = domain.NewQueue(queueName)
		r.queues[queueName] = queue
	}

	if len(queue.Messages) > 0 {
		message := queue.Messages[0]
		queue.Messages = queue.Messages[1:]
		r.mutex.Unlock()
		return message, nil
	}

	consumer := domain.NewQueueConsumer(timeout)
	queue.Consumers = append(queue.Consumers, consumer)
	r.mutex.Unlock()

	select {
	case message := <-consumer.ResponseChan:
		return message, nil
	case <-time.After(timeout):
		r.mutex.Lock()
		for i, c := range queue.Consumers {
			if c == consumer {
				queue.Consumers = append(queue.Consumers[:i], queue.Consumers[i+1:]...)
				break
			}
		}
		r.mutex.Unlock()
		return nil, errors.New("timeout")
	}
}