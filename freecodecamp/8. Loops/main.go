package main

import "fmt"

func bulkMessage(messageNumber int) float64 {
	totalCost := 0.0
	for i := 0; i < messageNumber; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

func maxMessage(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

func fizzbuzz() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	/*
		===== For Loop ======
	*/
	for i := 0; i < 10; i++ {
		fmt.Println("The number:", i)
	}

	// Infinity loop without condition
	// for j := 1; ; j++ {
	// 	fmt.Println("The number in infinity loop:", j)
	// }

	// Calculating message cost in for-loop
	for i := 1; i < 20; i += 5 {
		fmt.Printf("%v messages sent. And cost is: %.2f\n", i, bulkMessage(i))
	}

	// Calculating max message according to money in infinity for-loop
	for i := 1.0; i < 100; i += 7.5 {
		fmt.Printf("For $%.2f, %d message can be send.\n", i, maxMessage(i))
	}

	/*
		===== While Loop ======
	*/
	number := 1
	for number < 10 {
		fmt.Println("Number is increasing. Current number is:", number)
		number++
	}

	// FizzBuzz game
	fizzbuzz()
}
