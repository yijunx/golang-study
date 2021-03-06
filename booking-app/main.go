package main

import (
	"booking-app/helper"
	"fmt" // containers for functions
	"sync"
	"time"
)

var confName string = "cool conf" // syntax sugar, cannot use for const, and specify type
const confTickets uint = 50

var remainingTickets uint = confTickets
var bookings = make([]UserData, 0)

type UserData struct { // class that does not support inheritance
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

type AnyHow struct {
	user      UserData
	isChecked bool
}

var wg = sync.WaitGroup{}

// where to start the code??
func main() {
	greetUsers()

	// fmt.Printf("confName is type %T\n", confName)

	// fmt.Println("welcome to the", confName, "!")

	// var bookings [10]string // cannot mix type
	for true {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			// bookings[0] = firstName + "-" + lastName
			bookTicket(userTickets, firstName, lastName, email)
			// adding go there, makes it a GOROUTINE!!!!!
			wg.Add(1)                                    // add threads for the main to wait for
			go sendTicket(firstName, userTickets, email) // green thread, not OS level thread

			firstNames := getFirstNames()
			fmt.Printf("our firstnames of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end the programe
				fmt.Println("sold out BRUH")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("name not ok")
			}
		}

		// var noTicketsRemaining bool = remainingTickets == 0
	}

	wg.Wait()

	// checking on array
	// fmt.Printf("the whole array: %v\n", bookings)
	// fmt.Printf("the first element of the array: %v\n", bookings[0])
	// fmt.Printf("the type of the array: %T, length is %v\n", bookings, len(bookings))

	// checking on slice
	// fmt.Printf("the whole slice: %v\n", bookingSlice)
	// fmt.Printf("the first element of the array: %v\n", bookingSlice[0])
	// fmt.Printf("the type of the array: %T, length is %v\n", bookingSlice, len(bookingSlice))

}

func greetUsers() {
	fmt.Println("GREETwelcome to our conf:" + confName)
	fmt.Printf("we have total %v, and we have %v left\n", confTickets, remainingTickets)
	fmt.Println("get your tickets here..")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var name = booking.firstName
		firstNames = append(firstNames, name)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("your first name pls:")
	fmt.Scan(&firstName)
	fmt.Println("Your last name pls:")
	fmt.Scan(&lastName)
	fmt.Println("Your email pls:")
	fmt.Scan(&email)
	fmt.Println("number of tickets pls:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	// create a map for the user
	// create an empty map
	// var userData = make(map[string]string) // cannot mix datatypes
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets, // well the last one needs comma also
	}

	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thanks you! user %v booked %v tickets\n", firstName+lastName, userTickets)
	fmt.Printf("left with %v tickets\n", remainingTickets)
}

func sendTicket(firstName string, userTickets uint, email string) {
	time.Sleep(5 * time.Second) // just to simulate that this is slow..
	fmt.Println("##############")
	var ticket = fmt.Sprintf("%v tickets for %v", firstName, userTickets)
	fmt.Printf("sending ticket: %v \n Email: %v\n", ticket, email)
	fmt.Println("#######DONE###")
	wg.Done()
}
