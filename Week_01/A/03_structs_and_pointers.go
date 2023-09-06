package main

import "fmt"

// declare a person struct with two attributes
type person struct {
	name string
	age  int
}

func main() {
	// declare emptyP as a person but dont initialise
	var emptyP person
	fmt.Println("emptyP=", emptyP)

	// initialise an empty person struct and store it in variable p
	p2 := person{}
	fmt.Println("p2=", p2)

	// create a person struct with unnamed values
	p3 := person{"Bob", 20}
	fmt.Println(p3)

	// create a person struct with named values
	p4 := person{name: "Alice", age: 30}
	fmt.Println(p4)

	// access one attribute at a time
	fmt.Printf("Original p4.age=%d\n", p4.age)

	// set one attribute at a time
	p4.age = 51
	fmt.Printf("Updated p4.age=%d\n", p4.age)

	// Ampersand: "&p4" gets the memory address of p4
	// Asterix: "*ptrP4" gets the contents the address ptrP4 points to
	ptrP4 := &p4
	fmt.Printf("ptrP4=%v\n", ptrP4)
	fmt.Printf("ptrP4=%v\n", *ptrP4)

	// we can declare a pointer directly
	var emptyPtr *person
	// its value is nil if not set
	if emptyPtr == nil {
		fmt.Printf("emptyPtr=%v\n", emptyPtr)
	}

	// we can use pointers to hold addresses of arbitrary structs
	emptyPtr = &p4
	fmt.Printf("emptyPtr=%v\n", emptyPtr)

	// if we try and access an empty pointer we get a panic
	fmt.Println()
	fmt.Println("Please do not panic, whilst we panic....")
	fmt.Println("******************************************************")
	emptyPtr = nil
	fmt.Printf("emptyPtr=%v\n", *emptyPtr)

}
