package repository

import (
	domain "github.com/egor/watcher/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GetUser(username string) (domain.User, error)
}

type Target interface {
	Create(userId int, target domain.Target) (int, error)
	GetAll(userId int) ([]domain.Target, error)
	GetById(userId, targetId int) (domain.Target, error)
	Update(userId, targetId int, input domain.UpdateTargetInput) error
	Delete(userId, targetId int) error
	GetAllForWorker() ([]domain.Target, error)
	UpdateStatus(id int, status string) error
}

type Repository struct {
	Authorization
	Target
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Target:        NewDomainTargetPostgres(db),
	}
}
