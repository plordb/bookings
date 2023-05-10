package main

import (
	"encoding/json"
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
}
