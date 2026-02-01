package service

import (
	"errors"

	domain "github.com/egor/watcher/pkg/model"
	"github.com/egor/watcher/pkg/repository"
)

type DomainTargetService struct {
	repo repository.Target
}

func NewDomainTargetService(repo repository.Target) *DomainTargetService {
	return &DomainTargetService{repo: repo}
}

func validateUpdateInput(i domain.UpdateTargetInput) error {
	if i.Title == nil && i.URL == nil && i.Interval == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (s *DomainTargetService) Create(userId int, target domain.Target) (int, error) {
	return s.repo.Create(userId, target)
}
func (s *DomainTargetService) GetAll(userId int) ([]domain.Target, error) {
	return s.repo.GetAll(userId)
}
func (s *DomainTargetService) GetById(userId, targetId int) (domain.Target, error) {
	return s.repo.GetById(userId, targetId)
}

func (s *DomainTargetService) Update(userId, targetId int, input domain.UpdateTargetInput) error {

	if err := validateUpdateInput(input); err != nil {
		return err
	}
	return s.repo.Update(userId, targetId, input)
}
func (s *DomainTargetService) Delete(userId, targetId int) error {
	return s.repo.Delete(userId, targetId)
}
