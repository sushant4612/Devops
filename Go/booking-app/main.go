package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// package level variable
const conferenceTicket = 50

var conferenceName = "Go Conference"
var remainingTicket uint = 50
var bookings = make([]myData, 0)

type myData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	// var bookings [50]string       							 array syntax

	greetUser()

	//Ask for input
	firstName, lastName, email, userTicket := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTicket)
	if isValidEmail && isValidName && isValidTicketNumber {
		bookTickets(userTicket, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)
		//print first Names
		firstNames := FirstNames()
		fmt.Printf("These are the first name %v\n", firstNames)

		if remainingTicket == 0 {
			// end program
			fmt.Println("Our conference is boooked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ in it")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of ticket you book is invalid")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}

func FirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter you email id:")
	fmt.Scan(&email)

	fmt.Println("Enter number of ticket:")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTickets(userTicket uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTicket
	//create ma map

	var userData = myData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTicket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v ticket. You will recive confirmation on %v \n", firstName, lastName, userTicket, email)
	fmt.Printf("%v ticket remaining for %v\n", remainingTicket, conferenceName)
}

func sendTicket(userTicket uint, firstname string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v ticket for %v %v", userTicket, firstname, lastName)
	fmt.Println("###########################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###########################")
	wg.Done()
}
