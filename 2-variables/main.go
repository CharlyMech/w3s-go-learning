package main

import (
	"fmt"
)

/*
Contents from lessons (Go Variables chapter):
- Declare variables: https://www.w3schools.com/go/go_variables.php
- Declare multiple variables: https://www.w3schools.com/go/go_variable_multi.php
- Naming rules: https://www.w3schools.com/go/go_variable_naming_rules.php
*/

// 1- Declare variables:
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

// 2- Declare multiple variables
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

// 3- Naming rules
func namingRules() {
	// A variable name must start with a letter or an underscore character (_)
	fmt.Println("A variable name must start with a letter or an underscore character (_):")
	fmt.Println("\t- ✅ '_x' ; ✅ 'x'")
	fmt.Println("\t- ❌ '1x' ; ❌ '-x'")

	fmt.Println("")

	// A variable name cannot start with a digit
	fmt.Println("A variable name cannot start with a digit:")
	fmt.Println("\t- ✅ 'x' ; ❌ '1x'")

	fmt.Println("")

	// A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
	fmt.Println("A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ ):")
	fmt.Println("\t- ✅ 'my_var' ; ✅ 'var2' ; ✅ 'myVar'")
	fmt.Println("\t- ❌ 'my-var' ; ❌ 'var!'")

	fmt.Println("")

	// Variable names are case-sensitive (age, Age and AGE are three different variables)
	fmt.Println("Variable names are case-sensitive (age, Age and AGE are three different variables):")
	fmt.Println("\t 'x' and 'X' variable names, are not the same!")

	fmt.Println("")

	// There is no limit on the length of the variable name
	fmt.Println("There is no limit on the length of the variable name:")
	fmt.Println("\t- ✅ 'x' is a valid variable name!")
	fmt.Println("\t- ✅ 'abdcdefghijklmnopqrstuvwxyzABCDEFG...' is a valid variable name too!")

	fmt.Println("")

	// A variable name cannot contain spaces
	fmt.Println("A variable name cannot contain spaces:")
	fmt.Println("\t- ✅ 'myVar' and 'my_var' are valid variable names!")
	fmt.Println("\t- ❌ 'my var' is not a valid variable names!")

	fmt.Println("")

	// The variable name cannot be any Go keywords
	fmt.Println("The variable name cannot be any Go keywords:")
	fmt.Println("\t- ✅ 'myVar' ; ✅ 'typeName' ; ✅ 'funcName'")
	fmt.Println("\t- ❌ 'var' ; ❌ 'type' ; ❌ 'func' ; ❌ 'if' ; ❌ 'for' ; ❌ 'return'")

	fmt.Println("")

	// Variable naming types
	fmt.Println("Go variables can be named in different ways:")
	fmt.Println("\t- Snake case (not in common Go convention): 'snake_case'.")
	fmt.Println("\t- Camel case (not exported vars or functions): 'camelCase'.")
	fmt.Println("\t- Pascal case (exported vars or functions): 'PascalCase'.")

}

func main() {
	fmt.Println("- Variable declarations and types -")
	declareVariables()

	fmt.Println("")

	fmt.Println("- Multiple variable declaration -")
	multipleDeclarations()

	fmt.Println("")

	fmt.Println("- Naming rules -")
	namingRules()
}
