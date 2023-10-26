package handlers

import (
	"fmt"
	"myapp/data"
	"net/http"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/macketus/niibigo"
)

// Handler is the type of handlers, and gives access to NiibiGo and models
type Handlers struct {
	App    *niibigo.Niibigo
	Models data.Models
}

// Home is the handler to render the homepage
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

// Gopage is the handler to demonstrate rendering a Go template
func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.GoPage(w, r, "home", nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.JetPage(w, r, "jet-template", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	myData := "bar"

	h.App.Session.Put(r.Context(), "foo", myData)

	myValue := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", myValue)

	err := h.App.Render.JetPage(w, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) JSON(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID      int64    `json:"id"`
		Name    string   `json:"name"`
		Hobbies []string `json:"hobbies"`
	}

	payload.ID = 10
	payload.Name = "Ian Tetteh"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteJSON(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) XML(w http.ResponseWriter, r *http.Request) {
	type Payload struct {
		ID      int64    `xml:"id"`
		Name    string   `xml:"name"`
		Hobbies []string `xml:"hobbies>hobby"`
	}

	var payload Payload
	payload.ID = 10
	payload.Name = "Dela Seade"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteXML(w, http.StatusOK, payload)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

// DownloadFile is the handler to demonstrate file download
func (h *Handlers) DownloadFile(w http.ResponseWriter, r *http.Request) {
	h.App.DownloadFile(w, r, "./public/images", "celeritas.jpg")
}

func (h *Handlers) TestCrypto(w http.ResponseWriter, r *http.Request) {
	plainText := "Hello World"
	fmt.Fprint(w, "Unencrypted:"+plainText+"\n")
	encrypted, err := h.encrypt(plainText)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Error500(w, r)
		return
	}
	fmt.Fprint(w, "Encrypted: "+encrypted+"\n")

	decrypted, err := h.decrypt(encrypted)
	if err != nil {
		h.App.ErrorLog.Println(err)
		h.App.Error500(w, r)
		return
	}
	fmt.Fprint(w, "Decrypted: "+decrypted+"\n")
}