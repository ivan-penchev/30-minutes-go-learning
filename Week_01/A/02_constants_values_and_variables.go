// Every go file MUST start with a package as its first line.

// Package main is special, because it can compile into a binary
package main

import (
	"fmt"
	"strconv"
)

// *********************************************************************************************************
// anything outside of a function is global to the package
// *********************************************************************************************************

// single declaration for constants and variables
// // constants have a fixed value that cannot be changed after they are defined.
// // they are evaluated at compile-time and must be initialized with a constant value (e.g., literals or previously defined constants).
const DEBUG = false

// // variables can have their values modified throughout the execution of the program.
// // they can be initialized with a value, but if not initialized, they get a default zero value based on their type (e.g., 0 for numeric types, "" for strings, nil for pointers, etc.).
var sessionEnabled = true

// multiple declaration for constants
const (
	CONSTANT_A = "11"
	CONSTANT_B = "22"
	CONSTANT_C = 45
)

// multiple declaration for variables, either initialise or just declare
var (
	i, j = 1, 2
	k    string
)

// Package main has an entry point function, called main()
func main() {

	// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	// The format 'verbs' are derived from C's but are simpler.
	// more details can be found in the go doc https://pkg.go.dev/fmt

	fmt.Println() // for spacing

	// *********************************************************************************************************
	// values
	// *********************************************************************************************************

	// backtick is used for long strings that might wrap lines
	// it helps here to wrap the double quotes
	fmt.Println(`string arithmetic: "go" + "lang"=`, "go"+"lang")

	fmt.Println("print integer as a string with Println: 1+1=", 1+1)
	fmt.Printf("print integer as an integer with Printf:  1+1=%d\n", 1+1)
	fmt.Println("print float as a string with println: 7.0/3.0=", 7.0/3.0)
	fmt.Printf("print float as a float with printf: 7.0/3.0=%f\n", 7.0/3.0)
	fmt.Printf("print float as float with 2 decimal places: 7.0/3.0=%.2f\n", 7.0/3.0)
	fmt.Println("print boolean as a string: true AND false=", true && false)
	fmt.Printf("print boolean as a boolean: true OR false=%t\n", true || false)
	fmt.Printf("print boolean as a boolean: NOT true=%t\n", !true)

	// *********************************************************************************************************
	// constants, as stated, cannot be changed
	// *********************************************************************************************************
	fmt.Printf("constants A=%s, B=%s, C=%d, debug=%t\n", CONSTANT_A, CONSTANT_B, CONSTANT_C, DEBUG)

	// *********************************************************************************************************
	// variables allow us to store and manipulate values
	// *********************************************************************************************************

	// declare and initialise with `:=` (Short Variable Declaration)
	// // the := operator allows you to declare and initialize a new variable in a concise way, inferring the variable's type from the assigned value.
	// // it can only be used within a function body (not at the package level).
	a1 := "I've been declared and initialised in a single line"
	fmt.Printf("String a1=%s\n", a1)

	// declare with var and initialise/replace with `=` (Variable Assignment)
	// // the = operator is used for variable assignment, which assigns a value to an already declared variable.
	// // it is used when you want to change the value of an existing variable. And can be used at at the package level.
	var a2 string
	fmt.Printf("String uninitialised a2=%s\n", a2)
	a2 = "I've been declared and initialised"
	fmt.Printf("String initialised a2=%s\n", a2)

	// multiple assignment in the same declare statement
	var b, c int = 1, 2
	fmt.Printf("Multiple assignment b=%d c=%d\n", b, c)

	var e int
	fmt.Printf("Uninitialsed e=%d\n", e)

	// boolean assignment
	var d = true
	fmt.Printf("Boolean d=%t\n", d)

	// *********************************************************************************************************
	// complex example
	// *********************************************************************************************************

	// set two variables
	anInt := 1
	aString := "2"

	// convert string to int, it might go wrong so we have to check the error
	aStringAsInt, err := strconv.Atoi(aString)

	if err == nil {
		combinedInt := anInt + aStringAsInt

		// fmt.Sprintf returns a formatted string
		combinedString := fmt.Sprintf("%d", anInt) + aString

		// print the result
		fmt.Printf("combinedInt: 1+2=%d, combinedString: 1+2=%s\n\n", combinedInt, combinedString)
	}

}
