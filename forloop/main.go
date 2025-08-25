package main

import "fmt"

func main() {
	//For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	cnt := 0
	for {
		fmt.Println("Infinite Loop")
		cnt++
		if cnt == 3 {
			break
		}
	}
}
