package main

import (
	"errors"
	"fmt"
)

// Error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
		// return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, _ := divide(10, 0)
	fmt.Println("Result : ", result)
}
