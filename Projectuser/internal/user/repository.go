package user

import (
	"context"
	"go-fundamental/projectGO/Projectuser/internal/domain"
	"log"
)

type DB struct {
	Users []domain.User
	MaxID uint64
}

type (
	Repository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
	}

	repo struct {
		db  DB
		log *log.Logger
	}
)

func NewRepo(db DB, log *log.Logger) Repository {
	return &repo{
		db:  db,
		log: log,
	}
}

func (r *repo) Create(ctx context.Context, user *domain.User) error {

	r.db.MaxID++
	user.ID = r.db.MaxID
	r.db.Users = append(r.db.Users, *user)
	r.log.Println("repository created")
	return nil
}

func (r *repo) GetAll(ctx context.Context) ([]domain.User, error) {
	return r.db.Users, nil
}
