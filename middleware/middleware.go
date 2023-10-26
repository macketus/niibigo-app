package middleware

import (
	"myapp/data"

	"github.com/macketus/niibigo"
)

type Middleware struct {
	App    *niibigo.Niibigo
	Models data.Models
}
