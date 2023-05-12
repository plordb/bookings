package main

import (
	"errors"
	"fmt"
	"net/http"
)

// 03-04

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page")
}

func About(w http.ResponseWriter, r *http.Request) {

	sum := AddValues(2, 3)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 3 is %d", sum))
}

func AddValues(x int, y int) int {

	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {

	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")

		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func divideValues(x float32, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("cannot divide by zero")

		return 0, err
	}

	result := x / y

	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
