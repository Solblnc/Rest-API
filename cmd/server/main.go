package main

import (
	"fmt"
	"github.com/Solblnc/Rest-API/internal/coment"
	"github.com/Solblnc/Rest-API/internal/database"
	transport "github.com/Solblnc/Rest-API/internal/transport/http"
	"log"
	"net/http"
)

// struct for the pointers

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Server is running")

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
		log.Fatal(err)
	}
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
