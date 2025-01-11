package main

import "fmt"

// Simple struct (public properties)
type Car struct {
	Name  string
	Model string
	Color string
}

// Ambedded struct
type Truck struct {
	wheelNumber int
	Car
}

func main() {
	// Simple struct
	malibu := Car{Name: "Malibu 2", Model: "Chevrolet", Color: "black"}

	spark := Car{}
	spark.Name = "Spark 3"
	spark.Model = "Chevrolet"
	spark.Color = "White"

	fmt.Printf("%s's car %s color is %s\n", malibu.Model, malibu.Name, malibu.Color)
	fmt.Printf("%s's car %s color is %s\n", spark.Model, spark.Name, spark.Color)

	// Anonymous struct
	person := struct {
		Name   string
		Age    int
		Gender string
		isDied bool
	}{
		Name: "A'zam",
		Age:  21,
	}

	fmt.Printf("My name is %s and I am %d years old.\n", person.Name, person.Age)

	// Embedded struct
	truck := Truck{
		wheelNumber: 8,
		Car: Car{
			Name:  "Volvo FMX",
			Model: "FMX 540 8x4 Heavy Duty Tipper",
			Color: "white, blue and red",
		},
	}

	fmt.Printf("The truck of %s, model %s has %d wheels and %s colors.\n", truck.Name, truck.Model, truck.wheelNumber, truck.Color)
}
