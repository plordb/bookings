package main

import "fmt"

// 02-09

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 12,
	}

	PrintInfo(gorilla)

}

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (d Gorilla) Says() string {
	return "grunt"
}

func (d Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")

}
