package main

import (
	"fmt"
)

/*
Contents from lessons (Go Output):
- Output Functions: https://www.w3schools.com/go/go_output.php
- Formatting Verbs:
*/

var I, J string = "Hello", "World"
var X, Y = 10, 20

func output() {
	/*
		Three types of output functions in Go:
		1. Print() - prints its arguments with their default format.
		2. Println() - similar to 'Print()' with the difference that a whitespace is added between the arguments, and a newline is added at the end.
		3. Printf() - first formats its argument based on the given formatting verb and then prints them.
			(Generals)
			- `%v` - default format
			- `%#v` - Go-syntax representation of the value.
			- `%T` - type of the value.
			- `%%` - a literal percent sign; consumes no value.

			(Boolean)
			- `%t` - the word true or false.

			(Integer)
			- `%b` - base 2.
			- `%c` - the character represented by the corresponding Unicode code point.
			- `%d` - base 10.
			- `%o` - base 8.
			- `%O` - base 8 with 0o prefix.
			- `%q` - a single-quoted rune literal safely escaped with Go syntax.
			- `%x` - base 16, with lower-case letters for a-f.
			- `%X` - base 16, with upper-case letters for A-F.
			- `%U` - Unicode format: U+1234; same as "U+%04X".

			(Floating-point and complex constituents)
			- `%b` - decimalless scientific notation with exponent a power of two,
						in the manner of strconv.FormatFloat with the 'b' format,
						e.g. -123456p-78
			- `%e` -	scientific notation, e.g. -1.234456e+78
			- `%E` -	scientific notation, e.g. -1.234456E+78
			- `%f` -	decimal point but no exponent, e.g. 123.456
			- `%F` -	synonym for %f
			- `%g` -	%e for large exponents, %f otherwise. Precision is discussed below.
			- `%G` -	%E for large exponents, %F otherwise
			- `%x` -	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
			- `%X` -	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20

			The exponent is always a decimal integer.
			For formats other than %b the exponent is at least two digits.

			(String and slice of bytes)
			- `%s` - the uninterpreted bytes of the string or slice
			- `%q` - a double-quoted string safely escaped with Go syntax
			- `%x` - base 16, lower-case, two characters per byte
			- `%X` - base 16, upper-case, two characters per byte

			(Slice)
			- `%p` - address of 0th element in base 16 notation, with leading 0x

		(etc with Pointer, default format for `%v`, Compound objects, etc.)

		DOCS from `fmt` package: https://pkg.go.dev/fmt
	*/

	fmt.Println("- The use of 'Print()' function -")
	fmt.Print(I)
	fmt.Print(J)
	fmt.Println("\n(+ use of '\\n' at the end to skip line)")
	fmt.Print(I, "\n")
	fmt.Print(J, "\n")
	fmt.Print(I, "\n", J, "\n")
	fmt.Print(I, " ", J, "\n")
	fmt.Println("You can also use 'Print()' to print numbers:")
	fmt.Print(X, Y, "\n")

	fmt.Println("")
	fmt.Println("- The use of 'Println()' function -")
	fmt.Println(I, J)

	fmt.Println("")
	fmt.Println("- The use of 'Printf()' function -")
	fmt.Printf("I has value: %v and type: %T\n", I, I)
	fmt.Printf("X has value: %v and type: %T", X, X)
}

func formattingVerbs() {
	var Z float64 = 15.5
	var K int = 12
	var TXT string = "Hello World!"
	var BOOLEAN_T bool = true
	var BOOLEAN_F bool = false

	fmt.Println("- General formatting:")
	fmt.Printf("%v\n", Z)
	fmt.Printf("%#v\n", Z)
	fmt.Printf("%v%%\n", Z)
	fmt.Printf("%T\n", Z)

	fmt.Printf("%v\n", TXT)
	fmt.Printf("%#v\n", TXT)
	fmt.Printf("%T\n", TXT)

	fmt.Println("- Integer formatting:")
	fmt.Printf("%b\n", K)
	fmt.Printf("%d\n", K)
	fmt.Printf("%+d\n", K) // Base 10 and always show sign
	fmt.Printf("%o\n", K)
	fmt.Printf("%O\n", K)
	fmt.Printf("%x\n", K)
	fmt.Printf("%X\n", K)
	fmt.Printf("%#x\n", K)  // Base 16, with leading 0x
	fmt.Printf("%4d\n", K)  // Pad with spaces (width 4, right justified)
	fmt.Printf("%-4d\n", K) // Pad with spaces (width 4, left justified)
	fmt.Printf("%04d\n", K) //Pad with zeroes (width 4)

	fmt.Println("- Floating-point formatting:")
	fmt.Printf("%e\n", Z)
	fmt.Printf("%f\n", Z)
	fmt.Printf("%.2f\n", Z)  // Default width, precision 2
	fmt.Printf("%6.2f\n", Z) // Width 6, precision 2
	fmt.Printf("%g\n", Z)

	fmt.Println("- String formatting:")
	fmt.Printf("%s\n", TXT)
	fmt.Printf("%q\n", TXT)
	fmt.Printf("%8s\n", TXT)  // Prints the value as plain string (width 8, right justified)
	fmt.Printf("%-8s\n", TXT) // Prints the value as plain string (width 8, left justified)
	fmt.Printf("%x\n", TXT)
	fmt.Printf("% x\n", TXT) // Prints the value as hex dump with spaces

	fmt.Println("- Boolean formatting:")
	fmt.Printf("%t\n", BOOLEAN_T)
	fmt.Printf("%t\n", BOOLEAN_F)
}

func main() {
	output()

	fmt.Print("\n\n")
	fmt.Println("- Formatting Verbs -")
	formattingVerbs()
}
