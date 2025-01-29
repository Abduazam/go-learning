package main

import "fmt"

type Triange struct {
	a int
	b int
	c int
}

func (trainge Triange) area() int {
	return (trainge.a + trainge.b + trainge.c) / 2
}

func main() {
	a := 6
	b := 8
	c := 10

	traingle := Triange{
		a: a,
		b: b,
		c: c,
	}

	fmt.Printf("Area of traingle (%d, %d, %d) is %d.\n", a, b, c, traingle.area())
}
