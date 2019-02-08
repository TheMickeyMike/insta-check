package executor

import (
	"github.com/TheMickeyMike/insta-check/pkg/service"
)

type Executor struct {
	usernamesQueue chan<- string
	resultsQueue   <-chan *result
}

func NewExecutor(workersCount, concurrentlyCount int, instagramService *service.Instagram) *Executor {
	usernamesQueue := make(chan string, concurrentlyCount)
	resultsQueue := make(chan *result, workersCount*concurrentlyCount)

	executor := &Executor{usernamesQueue, resultsQueue}

	for id := 1; id <= workersCount; id++ {
		go NewWorker(id, usernamesQueue, resultsQueue).Run(instagramService)
	}
	return executor
}

func (executor *Executor) RunTask(usernamesToCheck []string) <-chan *result {
	for _, username := range usernamesToCheck {
		executor.usernamesQueue <- username
	}
	close(executor.usernamesQueue)
	return executor.resultsQueue
}
