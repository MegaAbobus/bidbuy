package order

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
)

type Service interface {
	InsertOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	GetOrder(ctx context.Context, ID uint) (*entities.Order, error)
	FetchOrders(ctx context.Context) (*[]presenter.Order, error)
	UpdateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error)
	RemoveOrder(ctx context.Context, ID uint) error
}

type service struct {
	storage Storage
}

func NewService(s Storage) Service {
	return &service{
		storage: s,
	}
}

func (s *service) InsertOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	return s.storage.CreateOrder(ctx, order)
}

func (s *service) GetOrder(ctx context.Context, ID uint) (*entities.Order, error) {
	return s.storage.ReadOrder(ctx, ID)
}

func (s *service) FetchOrders(ctx context.Context) (*[]presenter.Order, error) {
	return s.storage.ListOrder(ctx)
}

func (s *service) UpdateOrder(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	return s.storage.UpdateOrder(ctx, order)
}

func (s *service) RemoveOrder(ctx context.Context, ID uint) error {
	return s.storage.DeleteOrder(ctx, ID)
}
