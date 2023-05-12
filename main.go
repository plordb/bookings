package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// 03-05

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.go.html")
}

func About(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "about.go.html")
}

func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("errr parsing template:", err)

		return
	}

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
