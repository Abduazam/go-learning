package main

import (
	"errors"
	"fmt"
)

// Simple error return function
func returnResult(number int) (int, error) {
	if number < 1 {
		return 0, fmt.Errorf("Number must be greater than 0")
	}

	return 5 * number, nil
}

// Own exception class
type DivisionByZeroException struct {
	divider int
}

func (e DivisionByZeroException) Error() string {
	return fmt.Sprintf("Can not divide %v by zero.", e.divider)
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, DivisionByZeroException{divider: a}
	}

	return a / b, nil
}

// Error package
func divide2(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No division by 0")
	}

	return a / b, nil
}

func main() {
	result, e := returnResult(0)
	if e != nil {
		fmt.Println("Error found:", e)
	} else {
		fmt.Println("Result:", result)
	}

	result2, e2 := divide(1, 0)
	if e2 != nil {
		fmt.Println(e2)
	} else {
		fmt.Println("Result:", result2)
	}

	result3, e3 := divide2(5, 0)
	if e3 != nil {
		fmt.Println(e3)
	} else {
		fmt.Println("Result:", result3)
	}
}
