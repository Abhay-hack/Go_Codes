package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Data Conversion
	var numi int = 42
	var data float64 = float64(numi)
	fmt.Println(data)
	str := strconv.Itoa(numi)
	fmt.Println(str)

	a := "1234"
	b, _ := strconv.Atoi(a)
	b = b + 10
	fmt.Println(b)
	c, _ := strconv.ParseFloat(a, 64)
	fmt.Println(c)

}
