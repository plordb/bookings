package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 02-13

type Person struct {
	FirstName string `json:"first_name"`
	LasttName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:has_dog`
}

func main() {
	myJson := `
[
	{
		"first_name": "Pablo",
		"Last_name": "Lorenzo",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Clark",
		"Last_name": "Kent",
		"hair_color": "black",
		"has_dog": false
	}
]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LasttName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LasttName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))
}
