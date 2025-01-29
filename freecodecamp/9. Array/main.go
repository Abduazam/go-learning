package main

import "fmt"

func calculateArrayLen(array []int) int {
	return len(array)
}

func calculateArrayCap(array []int) int {
	return cap(array)
}

func calculateMessaageCostByLength(message string) float64 {
	total := 0.0 + 0.01*float64(len(message))

	return total
}

func main() {
	/*
		===== Array ======
	*/
	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", array)

	// It fill other index with digit 0
	numbers := [5]int{1, 2}
	fmt.Println("Numbers:", numbers)

	// Initializing numbers into indexes
	digits := [5]int{1: 1, 4: 1}
	fmt.Println("Digits:", digits)

	// Size not defined array
	noSize := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("No size array:", noSize)

	/*
		===== Slice ======
	*/
	// Digit before : from index, itself included, digit after : till index, itself excluded
	slice := array[1:4]
	fmt.Println("Slice:", slice)

	// Calculating slice length - actual count of elements
	fmt.Println("Slice count is:", calculateArrayLen(slice))

	// Calculating slice cap - till how many elements can be in slice
	fmt.Println("Slice cap is:", calculateArrayCap(slice))

	// Create slice from full of array
	fullSlice := array[:]
	fmt.Println("Full slice:", fullSlice)

	// From started index of array
	fromArray := array[4:]
	fmt.Println("Slice from 4 index:", fromArray)

	// Till ended index of array
	tillArray := array[:2]
	fmt.Println("Slice till 2 index:", tillArray)

	// String slice
	users := []string{"Azam", "Bobur", "Charos", "Dilnoza", "Feruzbek"}
	fmt.Println("Users name is:", users)

	messages := []string{
		"Hello World!",
		"Lorem ipsum doler amet",
		"Бесплатная проверка текста на ошибки, уникальность и SEO-анализ. Нейросети для генерации текста и картинок повысят эффективность вашей работы с контентом.",
	}

	for i := 0; i < len(messages); i++ {
		fmt.Printf("%s costs %.2f\n", messages[i], calculateMessaageCostByLength(messages[i]))
	}

	// Creating slice by make() function
	makeSlice := make([]int, 5, 10)
	fmt.Println("Slice by make func:", makeSlice)

	// Change value in slice
	makeSlice[2] = 3
	fmt.Println("Changed slice:", makeSlice)

	// Adding value in slice
	makeSlice = append(makeSlice, 10, 11, 12)
	fmt.Println("Extended slice:", makeSlice)

	// Merging to slices
	mergedSlice := append(makeSlice, tillArray...)
	fmt.Println("Merged slice:", mergedSlice)

	// Iterating slice
	for index, item := range mergedSlice {
		fmt.Printf("Index %d contains the digit: %d\n", index, item)
	}
}
