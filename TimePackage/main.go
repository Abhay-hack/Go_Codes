package main

import (
	"fmt"
	"time"
)

func main() {
	// Time Package
	currentTime := time.Now()
	fmt.Println("Current Time is", currentTime)

	formatted := currentTime.Format("02-01-2006, Monday, 03:04 PM")
	fmt.Println("Formatted Time is ", formatted)

	layout_str := "02/01/2006"
	dateStr := "15/06/2026"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted Time:", formatted_time)

	// add 1 more day to our current time
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("new date time:", new_date)
	fmt.Println("Formatted new time: ", new_date.Format("2006/01/02 Monday 03:04 PM"))

	// defer keyword => basically used to execute something just before the main fuction stops
	fmt.Println("Abhay")
	defer fmt.Println("Gupta")
	fmt.Println("Kumar")
}
