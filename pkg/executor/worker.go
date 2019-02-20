package executor

import (
	"fmt"

	"github.com/TheMickeyMike/insta-check/pkg/service"
)

type result struct {
	username  string
	available bool
	err       error
}

func (r result) String() string {
	return fmt.Sprintf("Username: %-10s Available: %-8t Error: %v", r.username, r.available, r.err)
}

type worker struct {
	id             int
	usernamesQueue <-chan string
	resultsQueue   chan<- *result
}

func NewWorker(id int, usernamesQueue <-chan string, resultsQueue chan<- *result) *worker {
	return &worker{id, usernamesQueue, resultsQueue}
}

func (worker *worker) Run(instagramService *service.Instagram) {
	var r result
	for username := range worker.usernamesQueue {
		r.username = username
		r.available, r.err = instagramService.UsernameIsAvailable(username)
		worker.resultsQueue <- &r
	}
}
