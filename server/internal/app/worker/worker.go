package worker

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
)

type Service interface {
	InsertWorker(ctx context.Context, worker *entities.Worker) (*entities.Worker, error)
	GetWorker(ctx context.Context, ID uint) (*entities.Worker, error)
	FetchWorkers(ctx context.Context) (*[]presenter.Worker, error)
	UpdateWorker(ctx context.Context, order *entities.Worker) (*entities.Worker, error)
	RemoveWorker(ctx context.Context, ID uint) error
}

type service struct {
	storage Storage
}

func NewService(s Storage) Service {
	return &service{
		storage: s,
	}
}

func (s *service) InsertWorker(ctx context.Context, worker *entities.Worker) (*entities.Worker, error) {
	return s.storage.CreateWorker(ctx, worker)
}

func (s *service) GetWorker(ctx context.Context, ID uint) (*entities.Worker, error) {
	return s.storage.ReadWorker(ctx, ID)
}

func (s *service) FetchWorkers(ctx context.Context) (*[]presenter.Worker, error) {
	return s.storage.ListWorker(ctx)
}

func (s *service) UpdateWorker(ctx context.Context, order *entities.Worker) (*entities.Worker, error) {
	return s.storage.UpdateWorker(ctx, order)
}

func (s *service) RemoveWorker(ctx context.Context, ID uint) error {
	return s.storage.DeleteWorker(ctx, ID)
}
