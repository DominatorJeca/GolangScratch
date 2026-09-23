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
	fmt.Println(getStatus(5000))

}

// Challenge: Consider the following function that receives
// a code number as arguments and returns a string with an
// HTTP reason, ex: 200 -> OK, 404 -> NOT FOUND
// If the code is not found returns: "Code {code} does not exist"
func getStatus(code int) string {
	if code == 200 {
		return "OK"
	}
	if code == 404 {
		return "NOT FOUND"
	}

	return fmt.Sprintf("Code %d does not exist", code)
}
