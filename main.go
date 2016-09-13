package main

import (
	"github.com/gernest/utron"
	_ "github.com/gernest/utron-todo/controllers"
	_ "github.com/gernest/utron-todo/models"
)

func main() {
	utron.Run()
}
