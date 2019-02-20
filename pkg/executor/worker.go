package executor

import (
	"fmt"

	"github.com/TheMickeyMike/insta-check/pkg/service"
	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

type result struct {
	username  string
	available bool
	err       error
}

func (r result) String() string {
	if r.available {
		r.username = green(r.username)
	} else {
		r.username = red(r.username)
	}
	return fmt.Sprintf("Username: %-18s Available: %-8t Error: %v", r.username, r.available, r.err)
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
	for username := range worker.usernamesQueue {
		r := result{username: username}
		r.available, r.err = instagramService.UsernameIsAvailable(username)
		worker.resultsQueue <- &r
	}
}
