package main

import (
	"fmt"
	"math"
)

func main() {
	var name string = "Abhay"
	// var sal float32 = 25550.55
	const pi = 3.14 // we can't update it
	age := 25       // declare and intialize with it :=
	isDeveloper := true
	// %s => string , %d => integer , %t => boolean , %q => quoted string
	fmt.Printf("My name is %s and my age is %d, Developer:%t\n", name, age, isDeveloper)
	fmt.Println(math.Sqrt(16))

	// same like this variable we can apply it to function
	// var PublicVariable int = 54 // we can use this in other files as it is public (first letter capital)
	// var privateVariable int = 24 // we can't use it in other files as it is private (first letter small)

	//If-else
	if age > 18 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}
