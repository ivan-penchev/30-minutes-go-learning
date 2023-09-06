// Every go file MUST start with a package as its first line.

// Package main is special, because it can compile into a binary
package main

import "fmt"

// Package main has an entry point function, called main()
func main() {

	// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	// The format 'verbs' are derived from C's but are simpler.
	// more details can be found in the go doc https://pkg.go.dev/fmt

	fmt.Println("hello world")
	fmt.Printf("hello %s\n", "gophers")
	fmt.Println("hello world and " + fmt.Sprintf("hello %s", "gophers"))

}
