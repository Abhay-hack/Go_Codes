package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Learning Json")
	person := Person{Name: "Abhay", Age: 20, IsAdult: true}
	fmt.Println("Person data is: ", person)

	//convert person into json encoding (Marshallling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling ", err)
		return
	}
	fmt.Println("Person Data is ", string(jsonData))

	//Decoding (UnMarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error Unmarshalling ", err)
		return
	}
	fmt.Println("Person data after unmarshalling", personData)
}
