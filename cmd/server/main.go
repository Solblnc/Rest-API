package main

import (
	"fmt"
	transport "github.com/Solblnc/Rest-API/internal/transport/http"
	"log"
	"net/http"
)

// struct for the pointers

type App struct {
}

func (app *App) Run() error {
	fmt.Println("Server is running")
	handler := transport.NewHandler()
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
