package main

import "fmt"

/*
Contents of lesson: https://www.w3schools.com/go/go_constants.php
*/

/*
Constant rules:
	- Constant names follow the same naming rules as variables.
	- Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
	- Constants can be declared both inside and outside of a function
*/

const PI int = 3 // Pi Is Exactly 3 !!
const NUMBER = 1

// Block declared constatns
const (
	X int    = 1
	Y string = "Hello"
	Z bool   = true
)

func main() {
	// If a variable name has a fixed value, is a constat variable.
	// In Go, we can use the keyword 'const' + UPPERCASE name (convention) syntax to declare constants

	fmt.Printf("Number 'pi' is a constant value: %v (%T)", PI, PI)

	fmt.Println("")

	fmt.Printf("Variable 'PI' (%v) is a typed constant, but 'NUMBER' (%b) is not!.", PI, NUMBER)

	fmt.Println("")

	// A const value cannot be changed
	const B int = 2
	// B = 3 // This would cause a compile-time error -> cannot assign to B

	fmt.Println("These constants are declared in block:")
	fmt.Printf("\t- X: %v (%T)\n", X, X)
	fmt.Printf("\t- Y: %v (%T)\n", Y, Y)
	fmt.Printf("\t- Z: %v (%T)\n", Z, Z)
}
