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
	app.Model.AutoMigrateAll()
	app.AddController(c.NewNote)

	port := fmt.Sprintf(":%d", app.Config.Port)
	app.Log.Info("staring server on port", port)
	log.Fatal(http.ListenAndServe(port, app))
}
