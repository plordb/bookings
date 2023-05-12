package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)

	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("errr parsing template:", err)

		return
	}

}
