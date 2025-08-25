package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Web Request
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	fmt.Println("Web Development")
	if err != nil {
		fmt.Println("Error in getting Get response ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n", res)

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response ", err)
		return
	}

	fmt.Println("response: ", string(data))

}
