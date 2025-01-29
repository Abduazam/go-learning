package main

import "fmt"

func main() {
	// Only if
	if 5 < 6 {
		fmt.Println("5 is less than 6")
	}

	// If-else
	var a int = 5
	b := 6
	if a > b {
		fmt.Printf("%d is greater than %d", a, b)
	} else {
		fmt.Printf("%d is greater than %d", b, a)
	}

	// If-else if-else
	x := 18
	if x >= 18 {
		fmt.Println("You are a student")
	} else if x < 18 && x >= 7 {
		fmt.Println("You are a pupil")
	} else {
		fmt.Printf("You are still going to kindergarten")
	}

	// Switch
	day := 8

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}

	// Swith multiple
	day2 := 5

	switch day2 {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	// Condition on variables
	age := 18
	isStudent := age >= 18
	fmt.Printf("Value: %t", isStudent)
}
