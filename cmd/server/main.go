package main

import (
	"github.com/Solblnc/Rest-API/internal/coment"
	"github.com/Solblnc/Rest-API/internal/database"
	transport "github.com/Solblnc/Rest-API/internal/transport/http"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// App - contain information of an app
type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"AppName":    app.Name,
		"AppVersion": app.Version,
	}).Info("Setting up application")

	db, err := database.NewDataBase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := coment.NewService(db)

	handler := transport.NewHandler(commentService)
	handler.SetUpRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	app := App{
		Name:    "Commenting service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		log.Error("Failed to run app")
		log.Fatal(err)
	}
}
