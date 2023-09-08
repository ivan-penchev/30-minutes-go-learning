package main

import "fmt"

func main() {

	// ########################
	// arrays are fixed length
	// ########################
	fmt.Println()
	// specify an array of length 5 (index = 0..4)
	var a [5]int
	for i := 0; i < 5; i++ {
		a[i] = i
	}
	// %v is really useful as it prints any underlying complex variable
	fmt.Printf("a=%v\n", a)

	// ########################
	// range can be used to iterate an array
	// ########################
	fmt.Println()
	for i, v := range a {
		fmt.Printf("Array range: index %d has value %d\n", i, v)
	}

	// if we go beyond the limit we will get a panic
	// a[5] = 5
	// panic: runtime error: index out of range [5] with length 5

	// ########################
	// range does not mutate the original data
	// Because the range loop copies the values from the slice to a local variable;
	// updating the local variable will not affect the slice.
	// ########################
	fmt.Println()
	s := []int{1, 1, 1}
	expected := []int{2, 2, 2}
	for _, n := range s {
		n += 1
	}
	fmt.Printf("Tried to change array expected %v got %v \n", expected, s)

	// ########################
	// slices are also arrays
	// they are initialised like arrays
	// but can also be extended beyond that with append
	// ########################
	fmt.Println()
	b := make([]int, 5)
	for i := 0; i < 5; i++ {
		b[i] = i
	}
	fmt.Printf("b=%v\n", b)
	b = append(b, 5)
	fmt.Printf("b=%v\n", b)

	// ########################
	// range can be used to iterate a slice
	// ########################
	fmt.Println()
	for i, v := range b {
		fmt.Printf("Slice range: index %d has value %d\n", i, v)
	}

	// maps are key'd structures
	m := make(map[string]int)
	n := make(map[int]string)
	for i := 0; i < 5; i++ {
		m[fmt.Sprintf("%d", i)] = i
		n[i] = fmt.Sprintf("%d", i)
	}

	// ########################
	// range can be used to iterate a map
	// ########################
	fmt.Println()
	for k, v := range m {
		fmt.Printf("Map range: index %s has value %d\n", k, v)
	}

}
