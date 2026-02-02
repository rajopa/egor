package service

import (
	"log/slog"
	"net/http"
	"sync"
	"time"

	domain "github.com/egor/watcher/pkg/model"
	"github.com/egor/watcher/pkg/repository"
)

type WorkerService struct {
	repo   repository.Target
	logger *slog.Logger
}

func NewWorkerService(repo repository.Target, logger *slog.Logger) *WorkerService {
	return &WorkerService{repo: repo}
}

func (s *WorkerService) Start() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		s.logger.Info("[Worker] Starting check cycle...")
		s.checkAll()
	}
}

func (s *WorkerService) checkAll() {

	targets, err := s.repo.GetAllForWorker()
	if err != nil {
		s.logger.Info("[Worker] Error getting targets: ", "error", err)
		return
	}

	var wg sync.WaitGroup

	for _, t := range targets {
		wg.Add(1)

		go func(target domain.Target) {
			defer wg.Done()

			client := http.Client{
				Timeout: 5 * time.Second,
			}

			status := "active"

			resp, err := client.Get(target.URL)
			if err != nil || resp.StatusCode != http.StatusOK {
				status = "error"
			}
			if resp != nil {
				resp.Body.Close()
			}
			err = s.repo.UpdateStatus(target.Id, status)
			if err != nil {
				s.logger.Error("error updating status", "id", target.Id, "error", err)
			}
		}(t)
	}

	wg.Wait()
	s.logger.Info("[Worker] Check cycle finished")
}
