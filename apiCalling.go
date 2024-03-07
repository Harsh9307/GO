package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// URL of the API endpoint
	apiUrl := "https://fakestoreapi.com/products"

	respoonse, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error making GET request: ", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(respoonse.Body)
	if err != nil {
		fmt.Printf("Error reading response body : %s", err)
		return
	}

	// Print the response
	fmt.Println(string(body))
}
