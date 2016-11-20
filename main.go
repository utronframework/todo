package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	c "github.com/utronframework/todo/controllers"
	"github.com/utronframework/todo/models"
)

func main() {

	// Start the MVC App
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register Models
	app.Model.Register(&models.Todo{})

	// CReate Models tables if they dont exist yet
	app.Model.AutoMigrateAll()

	// Register Controller
	app.AddController(c.NewTodo)

	// Start the server
	port := fmt.Sprintf(":%d", app.Config.Port)
	log.Fatal(http.ListenAndServe(port, app))
}
