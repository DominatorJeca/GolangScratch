package main

import "fmt"

func main() {
	destination := "Lisbon"
	travelers := 4
	bookingConfirmed := true

	pricePerNight := 89.90
	nights := 3
	total := pricePerNight * float64(nights)
	perPerson := total / float64(travelers)
	fmt.Printf("Trip to %s for %d friends. Booked: %t\n", destination, travelers, bookingConfirmed)
	fmt.Printf("Total: $%.2f, each pays $%.2f\n", total, perPerson)

	budget := 70.00
	withinBudget := perPerson <= budget

	fmt.Printf("Within budget: %v\n", withinBudget)
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
