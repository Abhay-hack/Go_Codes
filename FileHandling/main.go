package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//File Handling
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file: ", err)
		return
	}
	defer file.Close()
	content := "Hello World by Abhay"
	byteswritten, errors := io.WriteString(file, content+"\n")
	fmt.Println("Byte written ", byteswritten)
	if errors != nil {
		fmt.Println("Error while writing file: ", errors)
		return
	}

	// Open file for reading
	filee, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while opening file: ", err)
		return
	}
	defer filee.Close()

	// create a buffer to read the file content
	buffer := make([]byte, 1024)
	//	Readthe file content into the buffer
	for {
		n, err := filee.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading file", err)
			return
		}

		//Process the read content
		fmt.Println(string(buffer[:n]))
	}

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}
	fmt.Println(string(data))
}
