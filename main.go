package main

import (
	"fmt"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func isEmailValid(email string) string {
	for {
		fmt.Println("Enter your Email")
		fmt.Scan(&email)
		_, err := mail.ParseAddress(email)
		if err == nil {
			return email
		} else {
			fmt.Println("Email is not valid")
		}
	}
}

func greetUsers(conferenceName string, remainingTickets uint) {
	fmt.Printf("Welcome in the booking service of %v\n", conferenceName)
	fmt.Printf("There are %v tickets left\n", remainingTickets)
}

func isNameSurnameValid(Name string, NameOrSurname string) string {
	for {
		fmt.Printf("Enter your %v\n", NameOrSurname)
		fmt.Scan(&Name)
		match, _ := regexp.MatchString("[0-9]", Name)
		if match {
			fmt.Printf("%v can't be numeric!\n", NameOrSurname)
		} else if len(Name) < 3 {
			fmt.Printf("%v must be at least 3 characters long!\n", NameOrSurname)
		} else {
			return Name
		}
	}
}

func isNumberValid(numOfTickets uint, remainingTickets uint) uint {
	for {
		fmt.Println("Print number of tickets")
		fmt.Scan(&numOfTickets)
		if numOfTickets == 0 {
			fmt.Println("Number is not valid, try again")
			time.Sleep(2 * time.Second)
		} else if numOfTickets > remainingTickets {
			fmt.Println("You can't order", numOfTickets, "tickets because only", remainingTickets, "left!")
		} else {
			return numOfTickets
		}
	}
}

func RecordUserData(userName string, lastName string, Email string, numOfTickets string) map[string]string {
	var userData = make(map[string]string)
	userData["Name"] = userName
	userData["Surname"] = lastName
	userData["Email"] = Email
	userData["NumberOfTickets"] = numOfTickets
	return userData
}

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	greetUsers(conferenceName, remainingTickets)

	var bookings = make([]map[string]string, 0)

	for {
		var userName string
		var lastName string
		var Email string
		var numOfTickets uint

		userName = isNameSurnameValid(userName, "Name")
		lastName = isNameSurnameValid(lastName, "Surname")
		numOfTickets = isNumberValid(numOfTickets, remainingTickets)
		Email = isEmailValid(Email)

		fmt.Println("You've ordered", numOfTickets, "tickets on the name of", userName,
			"confirmation will be sent on Email:", Email)

		remainingTickets = remainingTickets - numOfTickets

		strNumOfTickets := strconv.FormatUint(uint64(numOfTickets), 10)
		userData := RecordUserData(userName, lastName, Email, strNumOfTickets)
		bookings = append(bookings, userData)

		var noTicketRemaining bool = remainingTickets == 0
		if noTicketRemaining {
			fmt.Println("There are no tickets left!")
			break
		}

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
	fmt.Println(bookings)

}
