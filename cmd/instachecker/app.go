package main

import (
	"fmt"
	"log"

	"github.com/TheMickeyMike/insta-check/pkg/client"
	"github.com/TheMickeyMike/insta-check/pkg/service"
)

// App is backbone for application
type App struct {
	instagram *service.Instagram
}

// Initialize setup app
func (app *App) Initialize() {
	fmt.Printf("%-13s: %s\n", "App name", version)
	fmt.Printf("%-13s: %s\n", "App version", name)

	httpClient := client.NewTrickyHTTP()
	app.instagram = service.NewInstagram(httpClient)
}

// Run 3 2 1.. Let's go
func (app *App) Run() {
	fmt.Println("Let's Go! 🚀")
	username := "maciej"
	result, err := app.instagram.UsernameIsAvailable(username)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Username: %-18s Available: %t\n", username, result)
}
