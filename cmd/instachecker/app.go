package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TheMickeyMike/insta-check/pkg/client"
	"github.com/TheMickeyMike/insta-check/pkg/executor"
	"github.com/TheMickeyMike/insta-check/pkg/service"
)

// App is backbone for application
type App struct {
	executor *executor.Executor
}

// Initialize setup app
func (app *App) Initialize() {
	fmt.Printf("%-13s: %s\n", "App name", name)
	fmt.Printf("%-13s: %s\n", "App version", version)

	httpClient := client.NewTrickyHTTP()
	instagramService := service.NewInstagram(httpClient)
	app.executor = executor.NewExecutor(2, 10, instagramService)
}

// Run 3 2 1.. Let's go
func (app *App) Run() {
	fmt.Printf("\nLet's Go! 🚀\n\n")

	usernames := []string{"maciej", "domi", "hdasjfb"}
	progress := len(usernames)
	for result := range app.executor.RunTask(usernames) {
		log.Printf("Result: %s", result)
		progress--
		if progress == 0 {
			break
		}
	}
	os.Exit(0)
}
