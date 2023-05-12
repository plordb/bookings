package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.go.html")
}

func About(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "about.go.html")
}
