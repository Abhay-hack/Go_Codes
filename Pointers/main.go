package main

import "fmt"

func modifyValueByReference(a *int) {
	*a = *a + 20
}

func main() {
	// => Pointers
	var num int
	num = 2

	// var ptr *int
	// ptr = &num
	var pointer *int
	if pointer == nil {
		fmt.Println("Value not assigned")
	}
	ptr := &num

	// fmt.Println("Num has value",num)
	fmt.Println("Pointer has value", ptr)
	fmt.Println("Pointer has this data:", *ptr)

	val := 5
	modifyValueByReference(&val)
	fmt.Println("Value Contains :", val)
}
