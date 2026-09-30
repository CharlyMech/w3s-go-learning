package main

import (
	"fmt"
)

/*
Contents from lessons (Go Variables chapter):
- Declare variables: https://www.w3schools.com/go/go_variables.php
- Declare multiple variables: https://www.w3schools.com/go/go_variable_multi.php
*/

// Declare variables:
// ! outsiderVar := "This will lead to an error" -> Outsider variables cannot be use ":=", only inside functions
var noValueOutsider int // -> Global file variables

func declareVariables() {
	// var [name] [type] = [value]
	// var [name] := [value] -> data type inferred

	// Primitive types and initial value
	var integer int = 2
	var floatNumber = 2.5
	text := "Hello Go lang" // -> inferred
	var boolean bool = true

	// Without initial value
	var text2 string

	fmt.Printf("Integer value: %v (%T)\n", integer, integer)
	fmt.Printf("Decimal/Float value: %v (%T)\n", floatNumber, floatNumber)
	fmt.Printf("String value: %v (%T)\n", text, text)
	fmt.Printf("String value without inital value: %v (%T)\n", text2, text2)
	text2 = "Now I have a value"
	fmt.Printf("Same string with value set: %v (%T)\n", text2, text2)
	fmt.Printf("Boolean value: %v (%T)\n", boolean, boolean)

	// Set value for the outsider var
	noValueOutsider = 3
	fmt.Printf("Boolean value: %v (%T)\n", boolean, boolean)
}

func multipleDeclarations() {
	var a, b, c int = 1, 2, 3

	fmt.Println("These variables where declared in one line with type declaration:")
	fmt.Printf("\t- Variable 'a': %v\n", a)
	fmt.Printf("\t- Variable 'b': %v\n", b)
	fmt.Printf("\t- Variable 'c': %v\n", c)

	var x, y = 1, "Hello"
	i, j := 2.5, false

	fmt.Println("These variables where declared in one line without type declaration:")
	fmt.Printf("\t- Variable 'x': %v (%T)\n", x, x)
	fmt.Printf("\t- Variable 'y': %v (%T)\n", y, y)
	fmt.Printf("\t- Variable 'i': %v (%T)\n", i, i)
	fmt.Printf("\t- Variable 'j': %v (%T)\n", j, j)

	var (
		q int
		w float32 = 3.14
		e string  = "World"
	)
	fmt.Println("These variables where declared in var block:")
	fmt.Printf("\t- Variable 'q': %v (%T)\n", q, q)
	fmt.Printf("\t- Variable 'w': %v (%T)\n", w, w)
	fmt.Printf("\t- Variable 'e': %v (%T)\n", e, e)

}

func main() {
	fmt.Println("- Variable declarations and types -")
	declareVariables()

	fmt.Println("")

	fmt.Println("- Multiple variable declaration -")
	multipleDeclarations()
}
