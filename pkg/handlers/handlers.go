package handlers

import (
	"fmt"
	"net/http"

	"github.com/plordb/bookings/pkg/config"
	"github.com/plordb/bookings/pkg/models"
	"github.com/plordb/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// performs some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hola, otra vez."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, r, "about.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	render.RenderTemplate(w, r, "make-reservation.gohtml", &models.TemplateData{})
=======
	render.RenderTemplate(w, "make-reservation.gohtml", &models.TemplateData{})
>>>>>>> 8094d222403a1136b6c99482bc2601dd5d1bf781
}

// Generals render the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	render.RenderTemplate(w, r, "generals.gohtml", &models.TemplateData{})
=======
	render.RenderTemplate(w, "generals.gohtml", &models.TemplateData{})
>>>>>>> 8094d222403a1136b6c99482bc2601dd5d1bf781
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	render.RenderTemplate(w, r, "majors.gohtml", &models.TemplateData{})
=======
	render.RenderTemplate(w, "majors.gohtml", &models.TemplateData{})
>>>>>>> 8094d222403a1136b6c99482bc2601dd5d1bf781
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	render.RenderTemplate(w, r, "search-availability.gohtml", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
=======
	render.RenderTemplate(w, "search-availability.gohtml", &models.TemplateData{})
>>>>>>> 8094d222403a1136b6c99482bc2601dd5d1bf781
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	render.RenderTemplate(w, r, "contact.gohtml", &models.TemplateData{})
=======
	render.RenderTemplate(w, "contact.gohtml", &models.TemplateData{})
>>>>>>> 8094d222403a1136b6c99482bc2601dd5d1bf781
}
