package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	fmt.Println("Learning Crud")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting ", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response ", err)
		return
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error Reading: ", err)
	// 	return
	// }
	// fmt.Println("Data: ", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding ", err)
		return
	}
	fmt.Println("Todo: ", todo)
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Abhay Gupta",
		Completed: true,
	}

	// convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshaling: ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"

	// Send Post Request
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending Response ", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll((res.Body))
	fmt.Println("Response Status: ", res.Status)
	fmt.Println("Data: ", string(data))

}

func performUpdateRequest() {
	todo := Todo{
		UserID:    2312,
		Title:     "Abhay Gupta crypt",
		Completed: false,
	}

	// convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshaling: ", err)
		return
	}

	// convert json data to string
	jsonString := string(jsonData)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// create PUT Request
	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT Response ", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Sending Request ", err)
		return
	}
	defer res.Body.Close()
	data, _ := io.ReadAll((res.Body))
	fmt.Println("Response Status: ", res.Status)
	fmt.Println("Data: ", string(data))
}

func performDeleteRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"
	// create PUT Request
	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)
	if err != nil {
		fmt.Println("Error creating Delete Request: ", err)
		return
	}

	// Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Sending Request ", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Response Status: ", res.Status)
}

func main() {

	//CRUD

	// Get Method
	performGetRequest()

	// Post Method
	performPostRequest()

	//Update Method
	performUpdateRequest()

	//Delete Method
	performDeleteRequest()

}
