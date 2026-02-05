package service

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/egor/watcher/kafka"
	domain "github.com/egor/watcher/pkg/model"
	"github.com/egor/watcher/pkg/repository"
)

type WorkerService struct {
	repo   repository.Target
	logger *slog.Logger
	kafka  *kafka.Producer
}

func NewWorkerService(repo repository.Target, logger *slog.Logger, kafka *kafka.Producer) *WorkerService {
	return &WorkerService{repo: repo, logger: logger, kafka: kafka}
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
	jobs := make(chan domain.Target, len(targets))

	var wg sync.WaitGroup

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := http.Client{Timeout: 5 * time.Second}
			for target := range jobs {

				status := "active"

				resp, err := client.Get(target.URL)

				if err != nil {
					status = "error"
				} else {
					if resp.StatusCode != http.StatusOK {
						status = "error"
					}
					resp.Body.Close()
				}
				err = s.repo.UpdateStatus(target.Id, status)
				if err == nil {
					s.kafka.SendMessage(context.Background(), "target_updates", target)

				}
			}

		}()
	}
	for _, t := range targets {
		jobs <- t
	}
	close(jobs)

	wg.Wait()
	s.logger.Info("[Worker] Check cycle finished")
}
