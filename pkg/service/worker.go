package service

import (
	"log"
	"net/http"
	"sync"
	"time"

	domain "github.com/egor/watcher/pkg/model"
	"github.com/egor/watcher/pkg/repository"
	"github.com/sirupsen/logrus"
)

type WorkerService struct {
	repo repository.Target
}

func NewWorkerService(repo repository.Target) *WorkerService {
	return &WorkerService{repo: repo}
}

func (s *WorkerService) Start() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("[Worker] Starting check cycle...")
		s.checkAll()
	}
}

func (s *WorkerService) checkAll() {

	targets, err := s.repo.GetAllForWorker()
	if err != nil {
		log.Printf("[Worker] Error getting targets: %v", err)
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
				logrus.Printf("[Worker] Error updating status for ID %d: %v", target.Id, err)
			}
		}(t)
	}

	wg.Wait()
	log.Println("[Worker] Check cycle finished")
}
