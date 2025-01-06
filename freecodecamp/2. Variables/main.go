package main

import "fmt"

func main() {
	// Go has 2 ways of declaring types

	// 1. Declaring variable type manual
	// integers
	var number int = 1          // Default integer type, 32 or 64 bit integer.
	var number8 int8 = 127      // 8-bit integer, (-128 to 127)
	var number16 int16 = -32768 // 16-bit integer, (-32,768 to 32,767)
	var number32 int32 = 4      // 32-bit integer (-2,147,483,648 to 2,147,483,647)
	var number64 int64 = -5     // 64-bit integer (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)

	fmt.Printf("Numbers:\n1. Integer: %d\n2. Integer8: %d\n3. Integer16: %d\n4. Integer32: %d\n5. Integer64: %d\n", number, number8, number16, number32, number64)

	// unsigned integers (not negative integers)
	var uNumber uint = 1     // Default non-negative integer type (32-bit or 64-bit based on system)
	var uNumber8 uint8 = 3   // 8-bit unsigned integer (0 to 255)
	var uNumber16 uint16 = 3 // 16-bit unsigned integer (0 to 65,535)
	var uNumber32 uint32 = 4 // 32-bit unsigned integer (0 to 4,294,967,295)
	var uNumber64 uint64 = 5 // 64-bit unsigned integer (0 to 18,446,744,073,709,551,615)

	fmt.Println()
	fmt.Printf("Unsigned Numbers:\n1. Integer: %d\n2. Integer8: %d\n3. Integer16: %d\n4. Integer32: %d\n5. Integer64: %d\n", uNumber, uNumber8, uNumber16, uNumber32, uNumber64)

	// floats
	var float32 float32 = 2.2      // 32-bit float-point number
	var float64 float64 = 3.343423 // 64-bit float-point number

	fmt.Println()
	// you can specify how many digits should be represents after dot by adding .number after %
	fmt.Printf("Float Numbers:\n1. Float 32: %f\n2. Float 64: %.1f\n", float32, float64)

	// byte
	// byte is alias for uint8
	var byteAsNumber byte = 5
	var byteAsLetter byte = 'A'

	fmt.Println()
	fmt.Printf("Bytes:\n1. Number: %d\n2. Letter: %c\n3. Letter as number: %d\n", byteAsNumber, byteAsLetter, byteAsLetter)

	// rune
	// rune is alias for int32
	var runeAsNumber rune = 563
	var runeAsLetter rune = 'ẽ'

	fmt.Println()
	fmt.Printf("Runes:\n1. Number: %d\n2. Letter: %c\n3. Letter as number: %d\n", runeAsNumber, runeAsLetter, runeAsLetter)

	// complex
	// complex types are used to represent complex numbers and imaginary numbers
	var complex64 complex64 = 3 + 6i       // float32 + float32 and i (abbreviation of imaginary, the square root of -1)
	var complex128 complex128 = 5.5 + 7.5i // float64 + float64 and i

	fmt.Println()
	fmt.Printf("Complexes:\n1. Complex 64: %v\n2. Complext 128: %v\n", complex64, complex128)

	// string
	var message string = "Hello World"

	fmt.Println()
	fmt.Println("String: " + message)

	// boolean
	var isGo bool = true

	fmt.Println()
	fmt.Printf("Boolean: %v\n", isGo)

	// 2. Declaring variables type automatic
	name := "A'zam"
	age := 21
	isMarried := false

	fmt.Println()
	fmt.Printf("Hello! My name is %s. I'm %d years old. Am I married: %v\n", name, age, isMarried)
}
