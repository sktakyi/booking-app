package main

import "fmt"

// Variables
const conferenceTickets = 50

var remainingTickets uint = 50
var bookings []string // Slice to store bookings

func main() {
	// Conference details
	conferenceName := "Go Conference"

	// Main program loop
	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()

		// Validate user input
		isValid := validateInput(firstName, lastName, email, userTickets)

		if isValid {
			// Book tickets and print booking details
			bookTickets(firstName, lastName, email, userTickets, conferenceName)

			// Print all bookings
			fmt.Printf("Current bookings: %v\n", bookings)
		} else {
			fmt.Println("Invalid input. Please try again.")
		}
	}

	// End program when tickets are sold out
	fmt.Println("Sorry, the conference is fully booked!")
}

// Function to greet users
func greetUsers(conferenceName string) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

// Function to get user input
func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// Function to validate user input
func validateInput(firstName, lastName, email string, userTickets uint) bool {
	isValidName := len(firstName) > 0 && len(lastName) > 0
	isValidEmail := len(email) > 5
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName && isValidEmail && isValidTicketNumber
}

// Function to book tickets
func bookTickets(firstName, lastName, email string, userTickets uint, conferenceName string) {

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
