package main

import (
	"fmt"
	"net/mail"
	"strings"
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

	var bookings []string

	for {
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
		bookings = append(bookings, userName+" "+lastName)

		var response string
		fmt.Println("Want to continue booking? Print 'Yes' or 'No'")
		fmt.Scan(&response)
		response = strings.TrimSpace(strings.ToLower(response))

		if response == "no" || response == "n" {
			break
		} else if response == "yes" || response == "y" {
			continue
		} else {
			fmt.Println("Not understood, exiting....")
		}

	}

}
