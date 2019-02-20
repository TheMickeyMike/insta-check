package executor

import (
	"fmt"

	"github.com/TheMickeyMike/insta-check/pkg/service"
	"github.com/fatih/color"
)

type result struct {
	username  string
	available bool
	err       error
}

func (r result) String() string {
	var username string
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	if r.available {
		username = green(r.username)
	} else {
		username = red(r.username)
	}
	return fmt.Sprintf("Username: %-18s Available: %-8t Error: %v", username, r.available, r.err)
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
