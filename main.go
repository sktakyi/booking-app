package main

import (
	"fmt"
	"strings"
)

func main() {
	// Variables
	var conferenceName = "Go Conference"
	var remainingTickets uint = 50

	// Welcome message
	fmt.Printf("\nWelcome to %v booking system!\n", conferenceName)
	fmt.Printf("We have %d tickets available.\n", remainingTickets)

	// Collect user information
	fmt.Println("\nEnter your details to book tickets:")
	var firstName, lastName string
	var ticketsToBook uint

	// Validate First Name
	fmt.Print("First Name: ")
	fmt.Scan(&firstName)
	if strings.TrimSpace(firstName) == "" {
		fmt.Println("First name cannot be empty!")
		return
	}

	// Validate Last Name
	fmt.Print("Last Name: ")
	fmt.Scan(&lastName)
	if strings.TrimSpace(lastName) == "" {
		fmt.Println("Last name cannot be empty!")
		return
	}

	// Validate Number of Tickets
	fmt.Print("Tickets to Book: ")
	_, err := fmt.Scan(&ticketsToBook)
	if err != nil {
		fmt.Println("Please enter a valid number for tickets.")
		return
	}
	if ticketsToBook <= 0 {
		fmt.Println("Invalid number of tickets. Please enter a positive number.")
		return
	}

	// Tickets Validation
	if ticketsToBook <= remainingTickets {
		remainingTickets -= ticketsToBook
		fmt.Printf("\nThank you, %s %s! Your %d tickets are booked. %d tickets remaining.\n",
			firstName, lastName, ticketsToBook, remainingTickets)
	} else {
		fmt.Printf("\nSorry, only %d tickets are available.\n", remainingTickets)
	}

	// Goodbye message
	fmt.Println("Thanks for using our booking system!")
}
