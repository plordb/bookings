package main

import (
	"fmt"
	"net/http"

	"github.com/plordb/bookings/pkg/handlers"
)

// 03-08

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
