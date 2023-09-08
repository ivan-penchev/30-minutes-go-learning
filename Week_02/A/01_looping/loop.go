package main

import "fmt"

func main() {

	// ###################################################
	// count to 5
	// ###################################################
	fmt.Println()
	for i := 1; i <= 5; i++ {
		fmt.Println("Count to 5:", i)
	}

	// ###################################################
	// loop forever
	// ###################################################
	fmt.Println()
	for {
		fmt.Println("Looping forever")
		fmt.Println("that might get boring so lets break out the loop")
		break
	}
	fmt.Println("Finished looping forever")

	// ###################################################
	// loop until true
	// There are no while loops in go because you can do the same with for
	// ###################################################
	fmt.Println()
	b := true
	for b {
		fmt.Println("Looping while b is true")
		b = false
	}
	fmt.Println("b is no longer true")

}
