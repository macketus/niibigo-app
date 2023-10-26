package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/macketus/niibigo"
)

type application struct {
	App      *niibigo.Niibigo
	Handlers *handlers.Handlers
	Models   data.Models  
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()

}
