package service

import (
	"context"
	"log/slog"

	"github.com/egor/watcher/kafka"
	domain "github.com/egor/watcher/pkg/model"
	"github.com/egor/watcher/pkg/repository"
)

type Authorization interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Target interface {
	Create(userId int, target domain.Target) (int, error)
	GetAll(userId int) ([]domain.Target, error)
	GetById(userId, targetId int) (domain.Target, error)
	Update(userId, targetId int, input domain.UpdateTargetInput) error
	Delete(userId, targetId int) error
}

type Worker interface {
	Start(ctx context.Context)
}

type Service struct {
	Authorization
	Target
	Worker
}

func NewService(repos *repository.Repository, logger *slog.Logger, kafka *kafka.Producer) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, logger),
		Target:        NewDomainTargetService(repos.Target, logger),
		Worker:        NewWorkerService(repos.Target, logger, kafka),
	}

}
