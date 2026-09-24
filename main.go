package main

import "fmt"

func main() {
	requests := []string{
		"GET /coffees HTTP/1.1", // 0
		"GET /beans HTTP/1.1",   // 1
		"GET /tea HTTP/1.1",     // 2
	}

	//fmt.Println(requests[0])

	requests = append(requests, "Hey this is a new element") // 3

	requests[len(requests)-1] = "POST /coffees HTTP/1.1"

	// List them
	// Loop over the requests to read their content
	for index, value := range requests {
		fmt.Printf("Index: %v, Value: %v\n", index, value)

		// Decide what to do with them...
		handle(index, value)
	}

}

func handle(index int, value string) {

}
