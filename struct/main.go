package main

import (
	"fmt"
)

//Structs and Interfaces

// => Structs
type Person struct {
	Name string
	Age  int
}
type Contact struct {
	Email string
	Phone string
}
type Address struct {
	City  string
	State string
}
type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

// => Method
func (p Person) Greet() string {
	return "Hello " + p.Name
}

// =>Interface
type Greeter interface {
	Greet() string
}

func main() {
	// Structs and Interfaces
	fmt.Println("Struct & Interfaces")

	// Creating a struct instance
	p := Person{Name: "Abhay", Age: 20}
	fmt.Println(p.Name)
	fmt.Println(p)

	fmt.Println(p.Greet())

	var employee1 Employee
	employee1.Person_Details = Person{
		Name: "Abhay",
		Age:  20,
	}
	employee1.Person_Contact.Email = "xyz@gmail.com"
	employee1.Person_Contact.Phone = "9672737472"
	employee1.Person_Address = Address{
		City:  "Bilgram",
		State: "Uttar Pradesh",
	}

	fmt.Println("Employee1 : ", employee1)

	// 	Structs = like objects/classes → hold data.
	// Methods = functions attached to structs (like class methods).
	// Interfaces = define behavior → if a struct has the methods, it implements the interface automatically

	var g Greeter
	g = Person{Name: "Abhay", Age: 20}
	fmt.Println(g.Greet())

}
