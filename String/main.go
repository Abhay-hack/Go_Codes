package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	var age int
	fmt.Scanln(&name, &age)
	fmt.Printf("My name is %s and my age is %d\n", name, age)

	// Reading a whole line, here buffer is a temporary storage
	reader := bufio.NewReader(os.Stdin)
	namee, _ := reader.ReadString('\n')
	fmt.Println("Details", namee)

	// Strings
	data := "Apple, Orage, Banana, Apple"
	parts := strings.Split(data, ",")
	fmt.Println(parts)
	fmt.Println(strings.Count(data, "Apple"))
	data = "     APPLE     FRUIT"
	fmt.Println(strings.TrimSpace(data))

	str1 := "Abhay"
	str2 := "Gupta"
	fmt.Println(strings.Join([]string{str1, str2}, " "))
}
