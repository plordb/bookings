package main

import (
	"errors"
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

	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
	}

	log.Println("result of division is", result)
}

func divide(x float32, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New(("cannot divde by 0"))
	}

	result = x / y
	return result, errors.New("")
}
