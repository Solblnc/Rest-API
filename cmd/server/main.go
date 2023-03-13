package main

import "fmt"

// struct for the pointers

type App struct {
}

func (app *App) Run() error {
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
