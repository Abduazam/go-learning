package main

import (
	"fmt"
)

func main() {
	// There are different example of using variable with differenct styles.

	/*
		===== Interpolations ======
	*/
	name := "A'zam"
	var age int8 = 21
	isMarried := false

	// %v => for default format
	fmt.Printf("Hello! My name is %v, and I'm %v y.o. Am I married: %v\n", name, age, isMarried)

	// %T => to print varaible type
	fmt.Printf("Name is %T | Age is %T | Married is %T\n", name, age, isMarried)

	// %x => for hex formar
	// %X => for hex format in uppercase
	fmt.Printf("Hex: %x\n", 255)
	fmt.Printf("Hex: %X\n", 255)

	// %s => for string types
	fmt.Printf("My name is %s\n", name)

	// %d => for integer types
	fmt.Printf("I am %d years old\n", age)

	// %f => for float types
	fmt.Printf("My overall GPA at university is %f\n", 4.0)
	// .1 declares how many digits should be represented after . in number
	fmt.Printf("My overall GPA at university is %.1f\n", 4.00)

	// %t => for boolean types
	fmt.Printf("Is he married: %t\n", true)

	// %c => for character in byte or rune
	fmt.Printf("Letter: %c\n", 'A')

	// %p => for pointer, memory address of variable
	fmt.Printf("Pointer: %p\n", &age)

	/*
		===== Constructions ======
	*/
	message := fmt.Sprintf(
		"Hello World. I am learning %s programming language. It's my %d lesson. Am I good at GO: %t, and I think I made %.1f progress",
		"GO",
		2,
		false,
		2.5,
	)

	hisName := "Bob"
	message2 := "Hey, his name is " + hisName + ", and he is " + fmt.Sprint(32)

	fmt.Println(message)
	fmt.Println(message2)
}
