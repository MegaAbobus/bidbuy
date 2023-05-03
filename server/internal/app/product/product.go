package product

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
)

type Service interface {
	InsertProduct(ctx context.Context, user *entities.Product) (*entities.Product, error)
	GetProduct(ctx context.Context, ID uint) (*entities.Product, error)
	FetchProducts(ctx context.Context) (*[]presenter.Product, error)
	UpdateProduct(ctx context.Context, order *entities.Product) (*entities.Product, error)
	RemoveProduct(ctx context.Context, ID uint) error
}

type service struct {
	storage Storage
}

func NewService(s Storage) Service {
	return &service{
		storage: s,
	}
}

func (s *service) InsertProduct(ctx context.Context, user *entities.Product) (*entities.Product, error) {
	return s.storage.CreateProduct(ctx, user)
}

func (s *service) GetProduct(ctx context.Context, ID uint) (*entities.Product, error) {
	return s.storage.ReadProduct(ctx, ID)
}

func (s *service) FetchProducts(ctx context.Context) (*[]presenter.Product, error) {
	return s.storage.ListProduct(ctx)
}

func (s *service) UpdateProduct(ctx context.Context, order *entities.Product) (*entities.Product, error) {
	return s.storage.UpdateProduct(ctx, order)
}

func (s *service) RemoveProduct(ctx context.Context, ID uint) error {
	return s.storage.DeleteProduct(ctx, ID)
}
