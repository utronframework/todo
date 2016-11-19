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
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(app.View == nil)
	app.Model.Register(&models.Todo{})
	app.Model.AutoMigrateAll()
	app.AddController(c.NewTodo)
	port := fmt.Sprintf(":%d", app.Config.Port)
	log.Fatal(http.ListenAndServe(port, app))
}
