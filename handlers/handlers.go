package handlers

import (
	"myapp/data"
	"net/http"

	"github.com/macketus/niibigo"
)

// Handler is the type of handlers, and gives access to NiibiGo and models
type Handlers struct {
	App    *niibigo.Niibigo
	Models data.Models
}

// Home is the handler to render the homepage
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
