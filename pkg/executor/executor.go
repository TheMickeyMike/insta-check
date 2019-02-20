package executor

import (
	"sync"

	"github.com/TheMickeyMike/insta-check/pkg/service"
)

type Executor struct {
	usernamesQueue chan<- string
	resultsQueue   <-chan *result
}

func NewExecutor(workersCount, concurrent int, instagramService *service.Instagram) *Executor {
	usernamesQueue := make(chan string, concurrent)
	resultsQueue := make(chan *result, workersCount*concurrent)

	var wg sync.WaitGroup
	wg.Add(workersCount)

	executor := &Executor{usernamesQueue, resultsQueue}

	for id := 1; id <= workersCount; id++ {
		go func(id int) {
			NewWorker(id, usernamesQueue, resultsQueue).Run(instagramService)
			wg.Done()
		}(id)
	}
	go func() {
		wg.Wait()
		close(resultsQueue)
	}()
	return executor
}

func (executor *Executor) RunTask(usernamesToCheck []string) <-chan *result {
	go func() {
		for _, username := range usernamesToCheck {
			executor.usernamesQueue <- username
		}
		close(executor.usernamesQueue)
	}()
	return executor.resultsQueue
}
