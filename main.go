package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	"github.com/gernest/utron/controller"
	c "github.com/utronframework/todo/controllers"
	"github.com/utronframework/todo/models"
)

func main() {
	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}
	app.Model.Register(&models.Todo{})
	app.AddController(controller.GetCtrlFunc(&c.Todo{}))
	port := fmt.Sprintf(":%d", app.Config.Port)
	log.Fatal(http.ListenAndServe(port, app))
}
