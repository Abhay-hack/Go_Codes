package main

import "fmt"

// Functions
func add(a int, b int) int {
	return a + b
}

func main() {
	// Functions => we can't declare function inside main
	result := add(3, 4)
	fmt.Println(result)
}
