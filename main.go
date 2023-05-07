package main

import (
	"log"

	"github.com/plordb/bookings/helpers"
)

// 02-10

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	myVar.TypeNumber = 5

	log.Println(myVar.TypeName)
	log.Println(myVar.TypeNumber)
}
