package main

import "fmt" // containers for functions

// how do i know which module contains my function? google it

// where to start the code??
func main() {

	confName := "cool conf" // syntax sugar, cannot use for const, and specify type
	var confTickets = 50
	var remainingTickets = 50

	fmt.Println(confTickets)
	fmt.Println(&confTickets) // cannot work for const..

	fmt.Printf("confName is type %T\n", confName)

	fmt.Println("welcome to the", confName, "!")
	fmt.Printf("welcome to the %v\n", confName)
	fmt.Printf("we have total %v, and we have %v left", confTickets, remainingTickets)
	fmt.Println("get your tickets here..")

	var userName string
	var userTickets int
	// ask user for their name
	// fmt.Scan(userName)

	// userName = "tom"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets\n", userName, userTickets)

}
