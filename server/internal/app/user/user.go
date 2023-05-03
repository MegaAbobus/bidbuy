package user

import (
	"bidbuy/internal/entities"
	"bidbuy/internal/presenter"
	"context"
)

type Service interface {
	InsertUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUser(ctx context.Context, ID uint) (*entities.User, error)
	FetchUsers(ctx context.Context) (*[]presenter.User, error)
	UpdateUser(ctx context.Context, order *entities.User) (*entities.User, error)
	RemoveUser(ctx context.Context, ID uint) error
}

type service struct {
	storage Storage
}

func NewService(s Storage) Service {
	return &service{
		storage: s,
	}
}

func (s *service) InsertUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	return s.storage.CreateUser(ctx, user)
}

func (s *service) GetUser(ctx context.Context, ID uint) (*entities.User, error) {
	return s.storage.ReadUser(ctx, ID)
}

func (s *service) FetchUsers(ctx context.Context) (*[]presenter.User, error) {
	return s.storage.ListUser(ctx)
}

func (s *service) UpdateUser(ctx context.Context, order *entities.User) (*entities.User, error) {
	return s.storage.UpdateUser(ctx, order)
}

func (s *service) RemoveUser(ctx context.Context, ID uint) error {
	return s.storage.DeleteUser(ctx, ID)
}
