package main

import (
	"fmt"
	"net/mail"
	"time"
)

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome in the booking service of %v\n", conferenceName)
	fmt.Printf("There are %v tickets left\n", remainingTickets)

	var userName string
	var lastName string
	var Email string
	var numOfTickets uint

	fmt.Println("Enter your name")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	for {
		fmt.Println("Enter your Email")
		fmt.Scan(&Email)

		if isEmailValid(Email) {
			break

		} else {
			fmt.Println("Email is not valid")
		}
	}

	for {
		fmt.Println("Print number of tickets")
		fmt.Scan(&numOfTickets)

		if numOfTickets == 0 {
			fmt.Println("Number is not valid, try again")
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	remainingTickets = remainingTickets - numOfTickets
}
