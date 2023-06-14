// first thing to do: go mod init DirNAME
package main

import (
	"GoBookingSystem/helper"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// struct: lightweight class w/o functions
type UserInfo struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// waitgroup to wait side thread to finish before the main ends
// usage: wg.Add(NUMofTHREAD) before goroutinue
// wg.wait() before the main thread ends
// wg.done() after the side thread finishes
var waitgroup = sync.WaitGroup{}

func main() {
	// package level variables (all packages can access)
	const conferenceTickets = 50

	var remainingTicket uint = 50 // uint: only positive

	conferenceName := "Go Conference"

	// Printf: format print; %v: var, %T: type
	// fmt.Printf("ConferenceName is a %T type, ConferenceTicket is a %T type.\n", conferenceName, conferenceTickets)
	greetings(conferenceName, conferenceTickets, remainingTicket)

	// array (static: predefine size and CANNOT increase after)
	// var bookingArr [50]string
	// bookingArr[0] = "John"
	// fmt.Printf("The Array type is %T and the elements are %v. \n", bookingArr, bookingArr)

	// slice (dynamic: start with len0, increase as needed)
	// var booking []string
	booking := []string{} // alternative
	// booking = append(booking, "John")
	// fmt.Printf("The Slice type is %T and the elements are %v. \n", booking, booking)

	firstInitials := []string{}

	// map a.k.a dictionary
	// map[keyT]valueT
	var userData = make(map[string]string) // make() to initialize
	// userData["Key"] = "Value"

	// create struct
	var dummyStruct = UserInfo{
		firstName:       "Paul",
		lastName:        "Lu",
		email:           "plu@123.com",
		numberOfTickets: 7,
	}
	dummies := []UserInfo{} // struct slice
	dummies = append(dummies, dummyStruct)

	dummiesMap := make(map[string]UserInfo) // struct map
	dummiesMap["User1"] = dummyStruct

	// retrieve value from the struct
	// NAME.ATTR
	fmt.Printf("In the dummyStruct, firstName is %v, lastName is %v, email is %v, numberOfTickets is %v.\n", dummyStruct.firstName, dummyStruct.lastName, dummyStruct.email, dummyStruct.numberOfTickets)

	// infinite loop
	for {

		userName, userTicket := helper.GetUserInput()

		if userTicket > remainingTicket {
			fmt.Printf("We only have %v tickets, you can't booked %v tickets.\n", remainingTicket, userTicket)
			continue
		}

		remainingTicket, booking = handleTickets(remainingTicket, userTicket, userName, booking)

		// added to the map
		userData[userName] = strconv.FormatUint(uint64(userTicket), 10)

		// add 1 thread to the waitgroup
		waitgroup.Add(1)
		// add 'go' to create another thread
		go sendEmail(userName, userTicket)

		if remainingTicket == 0 {
			fmt.Println("All tickets have sold out, come back next year!")
			break
		}
	}
	// loop thru each element: for i,e := range arrName {}
	for _, name := range booking {
		firstInitial := strings.ToUpper(name[0:1])
		firstInitials = append(firstInitials, firstInitial)
	}
	fmt.Printf("There are all our booking initials: %v\n", firstInitials)
	fmt.Printf("There are all user data: %v\n", userData)

	waitgroup.Wait()
}

// func NAME (params) (returns)
func greetings(conferenceName string, conferenceTickets int, remainingTicket uint) {

	fmt.Printf("Welcome to %v Booking System.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are left.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your ticket to attend.")
}

func handleTickets(remainingTicket uint, userTicket uint, userName string, booking []string) (uint, []string) {

	remainingTicket -= userTicket

	fmt.Printf("Thank you %v for booking %v tickets. \n", userName, userTicket)
	fmt.Printf("There are still %v left \n", remainingTicket)

	booking = append(booking, userName)
	return remainingTicket, booking
}

// use multithread to avoid waiting time
func sendEmail(userName string, userTicket uint) {
	fmt.Println("Please Wait...")
	time.Sleep(10 * time.Second)

	confirmation := fmt.Sprintf("User %v has booked %v tickets", userName, userTicket)
	fmt.Println("##########################")
	fmt.Printf("Email: %v has been sent!\n", confirmation)
	fmt.Println("##########################")

	waitgroup.Done()
}
