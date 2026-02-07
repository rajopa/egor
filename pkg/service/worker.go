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

func (s *WorkerService) Start(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.logger.Info("[Worker] Stopped gracefully")
			return
		case <-ticker.C:
			s.checkAll(ctx)
		}
	}
}

func (s *WorkerService) checkAll(ctx context.Context) {

	targets, err := s.repo.GetAllForWorker()
	if err != nil {
		s.logger.Info("[Worker] Error getting targets: ", "error", err)
		return
	}
	jobs := make(chan domain.Target, len(targets))

	var wg sync.WaitGroup

	client := http.Client{Timeout: 5 * time.Second}
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for target := range jobs {

				req, err := http.NewRequestWithContext(ctx, "GET", target.URL, nil)
				if err != nil {
					continue
				}
				resp, err := client.Do(req)

				if ctx.Err() != nil {
					return
				}

				status := true

				if err != nil {
					status = false
				} else {
					if resp.StatusCode != http.StatusOK {
						status = false
					}
					resp.Body.Close()
				}
				target.Status = status
				err = s.repo.UpdateStatus(target.Id, status)
				if err == nil {
					s.kafka.SendMessage(ctx, "target_updates", target)
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
