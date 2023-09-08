package main

import "fmt"

func main() {

	a := 1
	b := 1
	c := 2

	fmt.Println()
	if a == b {
		fmt.Println("A=B is a true statement")
	}

	fmt.Println()
	if b != c {
		fmt.Println("B is not equal to C is a true statement")
	}

	// ###################################################
	// if then / else if then / else
	// range condition
	// ###################################################
	fmt.Println()
	var d int
	if d > 0 {
		fmt.Println("d is a positive number")
	} else if d < 0 {
		fmt.Println("d is a negative number")
	} else {
		fmt.Println("d is zero")
	}

	// ###################################################
	// if NOT this condition
	// ###################################################
	fmt.Println()
	e := false
	if !e {
		fmt.Println("e is not true")
	}

	if !(1 == 2) {
		fmt.Println("1 == 2 is false, and NOT that is true")
	}

	// ###################################################
	// conditions and loops
	// ###################################################
	fmt.Println()
	myCondition := true
	for {
		fmt.Println("Looping while b2 is true")

		// our first conditional
		// "if true" run the code in the brackets
		if myCondition {
			myCondition = false
			// continue restarts the loop and the code below will not be reached
			continue
		}

		fmt.Println("Checking if b2 is false")

		// our second conditional
		// "if not true" run the code in the brackets
		if !myCondition {
			break
		}
	}

	// ###################################################
	// switch statement
	// ###################################################
	fmt.Println()
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

}
