package main

import "fmt"

func main() {
	people := 4
	mealPrice := 18.75
	subTotal := float64(people) * mealPrice
	tax := subTotal * 0.10
	total := subTotal + tax

	fmt.Printf("Final bill: $%.2f\n", total)

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
