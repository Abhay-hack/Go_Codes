package main

import "fmt"

func main() {
	// Arrays
	fmt.Println("Arrays")
	// var arr [3]int= [3]int{1,2,3}
	numbers := []int{12, 21, 32, 41, 52}
	for index, value := range numbers {
		fmt.Printf("Index: %d , Value: %d\n", index, value)
	}

	for _, val := range numbers {
		fmt.Print(val)
		fmt.Print(" ")
	}
	// len() , cap() => capacity
}
