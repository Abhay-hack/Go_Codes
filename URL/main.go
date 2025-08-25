package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning Url")
	myUrl := "https://example.com:8080/search/products?id=123&category=books"
	parseUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("Can't parse url", err)
		return
	}
	fmt.Printf("Type of Url %T\n", parseUrl)
	fmt.Println("Scheme: ", parseUrl.Scheme)
	fmt.Println("Host: ", parseUrl.Host)
	fmt.Println("Path: ", parseUrl.Path)
	fmt.Println("RawQuery: ", parseUrl.RawQuery)

	// Modify the url components
	parseUrl.Path = "/newPath"
	parseUrl.RawQuery = "username*abhay"

	// construct a url from a url object
	newUrl := parseUrl.String()
	fmt.Println("New Url: ", newUrl)
}
