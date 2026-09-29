package main

import (
	"fmt"

	"rsc.io/quote"
)

/*
Contents from lessons:
- Syntax: https://www.w3schools.com/go/go_syntax.php
- Comments: https://www.w3schools.com/go/go_comments.php
*/

func main() {
	fmt.Println("Hello World!")

	// Call the quote package method to display pithy sayings
	fmt.Println(quote.Go())
}
