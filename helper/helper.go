package helper // exported package name

import "fmt"

// Global Variable can be shared across all packages
// MUST BE capitialized
var GlobalVar = "Global Variable"

// name of the exported function must be capitalized
func GetUserInput() (string, uint) {

	// local variable
	var userName string
	var userTicket uint
	for {
		fmt.Println("Please enter your first name: ")
		fmt.Scan(&userName) // use pointer when scan an input value
		fmt.Println("How many tickets would you like: ")
		fmt.Scan(&userTicket)

		// input validation
		validTicket := userTicket > 0
		validName := len(userName) > 0

		if !validName {
			fmt.Println("Invalid Name")
			continue
		}

		if !validTicket {
			fmt.Println("Invalid Ticket Number")
			continue
		}

		break
	}

	return userName, userTicket
}
