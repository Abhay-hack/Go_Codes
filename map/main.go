package main

import "fmt"

func main() {
	//Map (key,vaue)

	// make is used to create and initialize slices , map , channels
	// as we see make create empty map then after we added keys to it
	m := make(map[string]int)
	m["age"] = 25
	fmt.Println(m["age"])

	// map literal -> we directly created map with predefined values
	pearson := map[string]string{"name": "Abhay", "Developer": "True"}
	fmt.Println(pearson["Developer"])
	delete(pearson, "Developer")
	fmt.Println("Developer Deleted")
	fmt.Println(pearson)
	content, exists := pearson["Developer"]
	fmt.Println(content)
	fmt.Println(exists)
	for i, v := range pearson {
		fmt.Println("key is", i, "Data is", v)
	}
	// fmt.Println(pearson["Developer"])

}
