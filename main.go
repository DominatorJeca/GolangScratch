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

	// The loop is set up for you (the _ ignores the index). Fill in the body.
	for _, line := range requests {
		// TODO: if line starts with "GET", append it to gets
		if strings.HasPrefix(line, "GET") {
			gets = append(gets, line)
		}
	}

	fmt.Printf("%d of them are GETs\n", len(gets))
}
