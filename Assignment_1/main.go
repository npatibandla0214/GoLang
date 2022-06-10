package main

import (
	"fmt"

	"./app"
)

func main() {
	//fmt.Println("My favorite number is", rand.Intn(10))
	x := 20
	y := 3
	fmt.Printf("The Sum of %d and %d is %d", x, y, app.Add(x, y))
	fmt.Printf("\nThe difference between %d and %d is %d", x, y, app.Sub(x, y))
	fmt.Printf("\nThe product of %d and %d is %d", x, y, app.Multiplication(x, y))
	var quotient, remainder = app.Division(x, y)
	fmt.Printf("\nThe quotient of %d and %d is %d", x, y, quotient)
	fmt.Printf("\nThe remainder of %d and %d is %d\n", x, y, remainder)
}
