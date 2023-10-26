package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/macketus/niibigo"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	// init niibigo
	nii := &niibigo.Niibigo{}
	err = nii.New(path)
	if err != nil {
		log.Fatal(err)
	}
	nii.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: nii,
	}

	myHandlers := &handlers.Handlers{
		App: nii,
	}

	app := &application{
		App:        nii,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
