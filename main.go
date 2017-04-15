package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gernest/utron"
	c "./controllers"
	"./models"
)

func main() {

	app, err := utron.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	app.Model.Register(&models.Note{})
	app.Model.Register(&models.Notebook{})
	app.Model.AutoMigrateAll()
	app.AddController(c.NewNote)
	app.AddController(c.NewNotebook)

	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
