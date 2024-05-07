package user

import (
	"context"
	"go-fundamental/projectGO/Projectuser/internal/domain"
	"log"
)

type (
	Service interface {
		CreateUser(ctx context.Context, firstname, lastname, email string) (*domain.User, error)
		GetAllUsers(ctx context.Context) ([]domain.User, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) CreateUser(ctx context.Context, firstname, lastname, email string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	s.log.Println("service created")
	return user, nil

}

func (s service) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
