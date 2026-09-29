package main

import (
	"fmt"
)

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

func main() {
	fmt.Println("- Variable declarations and types -")
	declareVariables()
}
