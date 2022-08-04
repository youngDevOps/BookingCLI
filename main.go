// Go CLI learning file
// To run - go run main.go
// Comment with lines to understand application.
//FMT - Println (Prints lins of code), Printf (Prints line with variables declared after), %v (Variable), %T (Prints data type.)

package main // Declaring package, called main which is the common one.

import (
	"fmt"     //Importing packages using a list of supplied ones. fmt is common package.
	"strings" // Strings package
)

func main() { // Declaring a function called main
	//Welcome Message
	conferenceName := "Go Conference" // Declaring a variable with alternative syntax
	const conferenceTickets int = 50  // Declaring a constant (value dosen't change)
	var remainingTickets uint = 50    // Declaring a variable
	bookings := []string{}            //Slice as dosen't have a declared limit/value in the array

	fmt.Printf("Welcome to %v booking application\n", conferenceName)                                                          // Using FMT format package to add a placeholder for the variable and declare it later
	fmt.Printf("We have a total of %v tickets and %v are still avalable to purchase. \n", conferenceTickets, remainingTickets) // Using FMT format package to add a placeholder for the variable and declare it later
	fmt.Println("Get your tickets here to attend")                                                                             // Using Println to print text

	//Function Logic
	//For loop to continue running the applicatioj
	for {
		var firstName string // variable string
		var lastName string  // variable string
		var email string     // variable string
		var userTickets uint // integer string

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) // Pointer - User input passing the memory address to the variable

		fmt.Println("Enter your Last name: ")
		fmt.Scan(&lastName) // Pointer - User input passing the memory address to the variable

		fmt.Println("Enter your email: ")
		fmt.Scan(&email) // Pointer - User input passing the memory address to the variable

		fmt.Println("Enter number of tickets required: ")
		fmt.Scan(&userTickets) // Pointer - User input passing the memory address to the variable

		//User input validation

		isValidName := len(firstName) >= 2 && len(lastName) >= 2 //checking length of string input
		isValidEmail := strings.Contains(email, "@") //checking if the string contains a @ character and returns in a boolean format
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets //Checking user tickets

		//Remaining ticket logic

		if isValidName && isValidEmail && isValidTicketNumber { //Logic for if a user tries to book more tickets than available
			remainingTickets = remainingTickets - userTickets   // logic to calculate remaning tickets
			bookings = append(bookings, firstName+" "+lastName) // booking variable slice which uses append() which is a built in go function with the array, and variables wanted to pass in.

			fmt.Printf("The whole slice: %v\n", bookings)    //Printing the whole slice
			fmt.Printf("The first slice: %v\n", bookings[0]) // Printing the first value using [0], remeber first value is 0 not 1
			fmt.Printf("Slice type: %T\n", bookings)         //Printing the Slice type
			fmt.Printf("Slice length: %v\n", len(bookings))  // Printing the Slice length using len()

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email) // Printing out user inputed information
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)                                                                // Printing remaining ticktes using memory logic

			firstNames := []string{}
			for _, booking := range bookings { // Gives the index and element value
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames) //Printng bookings

			//Checking if tickets are sold out and weather to continue
			if remainingTickets == 0 {
				//End Program
				fmt.Println("Our conference is booked out. Come back next year!")
				break //ends program
			}
		} else {  //Letting user which input failed.
			if !isValidName {
				fmt.Println("First name or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address entered does not contain an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}
