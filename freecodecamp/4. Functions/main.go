package main

import "fmt"

// Simple function
func sayHello(name string) {
	fmt.Println("Hello", name)
}

func add(x int8, y int8) int8 {
	return x + y
}

// Same type arguments
func subtract(x, y int8) int8 {
	return x - y
}

// Function inside function
func doubleAdd(add func(a, b int8) int8, c int8) int8 {
	return add(5, 4) + c
}

// Multiple retuns
func numbers(a, b, c int) (int, int, int) {
	return a * a, b * b, c * c
}

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("Division by zero.")
	}

	return x / y, nil
}

// Return by variable name inside function
func multiply(x, y int) (result int) {
	result = x * y
	return
}

// Recursion
func increment(i int8) int8 {
	if i == 11 {
		return 0
	}

	fmt.Println("Number", i)
	return increment(i + 1)
}

// Variadic
func sum(numbers ...int) int {
	var total int = 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println("Addition: 5 + 3 =", add(5, 3))

	fmt.Println("Subtraction: 10 - 1 =", subtract(10, 1))

	fmt.Println("Double add: 10 + 5 + 9 =", doubleAdd(add, 10))

	x, y, z := numbers(2, 3, 4)
	fmt.Println("Squares:", x, y, z)

	w, err := divide(4, 0)
	if err != nil {
		fmt.Println("Division: 4 / 0 =", err)
	} else {
		fmt.Println("Division: 4 / 0 =", w)
	}

	fmt.Println("Multiply: 2 * 3 =", multiply(2, 3))

	increment(5)

	fmt.Println("Sum of 1, 2, 3, 4, 5 is equal to", sum(1, 2, 3, 4, 5))

	// Anonymous
	hi := func(name string) {
		fmt.Println("Hi", name)
	}
	hi("John")

	// Inline
	func(msg string) {
		fmt.Println(msg)
	}("It's inline function format!")
}
