package main

import "fmt"

func main() {
	//Switch
	day := "Monday"
	switch day {
	case "Monday", "Tuesday": //case "Monday":
		fmt.Println("Starting Phase")
	case "Thursday", "Wednesday":
		fmt.Println("Middle Phase")
	default:
		fmt.Println("No")
	}
}
