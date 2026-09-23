package main

import "fmt"

func main() {
	// fmt.Println("HTTP/1.1 200 OK")
	// fmt.Println("Content-Type: text/plain")
	// fmt.Println()
	// fmt.Println("Hello, World!")

	content := "Hello from the other side"

	// 	requestMessage := "HTTP/1.1 200 OK"

	// 	formatMessage := `HTTP/1.1 200 OK
	// Content-Type: text/plain

	// ` + content

	//fmt.Println(requestMessage)
	//fmt.Println(formatMessage)
	fmt.Printf(`HTTP/1.1 200 OK
Content-Type: text/plain

%s`, content)

	response := fmt.Sprintf(`HTTP/1.1 200 OK
Content-Type: text/plain

%s`, content)

	fmt.Println(response)
}
