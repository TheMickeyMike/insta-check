package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TheMickeyMike/insta-check/pkg/client"
	"github.com/TheMickeyMike/insta-check/pkg/config"
	"github.com/TheMickeyMike/insta-check/pkg/executor"
	"github.com/TheMickeyMike/insta-check/pkg/service"
)

// App is backbone for application
type App struct {
	config   *config.AppConfig
	executor *executor.Executor
}

// Initialize setup app
func (app *App) Initialize() {
	fmt.Printf("%-13s: %s\n", "App name", name)
	fmt.Printf("%-13s: %s\n", "App version", version)
	app.config = config.LoadConfig()

	httpClient := client.NewTrickyHTTP()
	instagramService := service.NewInstagram(app.config.Instagram, httpClient)
	app.executor = executor.NewExecutor(4, instagramService)
}

// Run 3 2 1.. Let's go
func (app *App) Run() {
	fmt.Printf("\nLet's Go! 🚀\n\n")

	usernames := []string{"jack", "moly", "hdasjfb", "bdfsfhvs"}
	for result := range app.executor.RunTask(usernames) {
		log.Printf("Result: %s", result)
	}

	os.Exit(0)
}
