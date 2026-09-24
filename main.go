package main

import (
	"fmt"
	"strings"
)

func main() {
	requests := make([]string, 0, 8)

	requests = append(requests, "GET /coffees HTTP/1.1")
	requests = append(requests, "BREW /tea HTTP/1.1")
	requests = append(requests, "GET /beans HTTP/1.1")
	requests = append(requests, "PING /status HTTP/1.1")
	requests = append(requests, "GET /tea HTTP/1.1")

	fmt.Printf("Logged %d requests (capacity %d)\n", len(requests), cap(requests))

	gets := make([]string, 0)

	for _, line := range requests {
		if strings.HasPrefix(line, "GET") {
			gets = append(gets, line)
		}
	}

	fmt.Printf("%d of them are GETs\n", len(gets))

	// Here's the pattern: delete index 3 ("PING /status HTTP/1.1").
	requests = append(requests[:3], requests[4:]...)

	// TODO: "BREW /tea HTTP/1.1" is still at index 1. Delete it the same way.
	requests = append(requests[:1], requests[2:]...)

	fmt.Printf("Cleaned log (%d): %v\n", len(requests), requests)
}

//requests[:i] is everything before i (up to, not including it).
//requests[i:] is everything after i.
