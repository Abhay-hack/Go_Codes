package main

import "fmt"

func main() {
	//Slice => dynamic arrays
	s := []int{1, 2, 3}
	s = append(s, 4, 5)
	fmt.Println(s[1:5])
	numbersi := make([]int, 3, 5) // length and capacity, slice doubles it capacity , if it exceeds its capacity
	fmt.Println("Slice", len(numbersi))
	fmt.Println("Capacity", cap(numbersi))

}
