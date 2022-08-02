// Go CLI learning file
// To run - go run main.go
// Comment with lines to understand application.
//FMT - Println (Prints lins of code), Printf (Prints line with variables declared after), %v (Variable), %T (Prints data type.)

package main // Declaring package, called main which is the common one.

import "fmt" //Importing packages using a list of supplied ones. fmt is common package.

func main() { // Declaring a function called main
	//Welcome Message
	conferenceName := "Go Conference" // Declaring a variable with alternative syntax
	const conferenceTickets int = 50  // Declaring a constant (value dosen't change)
	var remainingTickets uint = 50    // Declaring a variable

	fmt.Printf("conferenceTickets is %T, remainingTicket is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName) // Printing out data types

	fmt.Printf("Welcome to %v booking application\n", conferenceName)                                                          // Using FMT format package to add a placeholder for the variable and declare it later
	fmt.Printf("We have a total of %v tickets and %v are still avalable to purchase. \n", conferenceTickets, remainingTickets) // Using FMT format package to add a placeholder for the variable and declare it later
	fmt.Println("Get your tickets here to attend")                                                                             // Using Println to print text
	//Function Logic
	var firstName string // variable string
	var lastName string  // variable string
	var email string     // variable string
	var userTickets int  // integer string

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // Pointer - User input passing the memory address to the variable

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName) // Pointer - User input passing the memory address to the variable

	fmt.Println("Enter your email: ")
	fmt.Scan(&email) // Pointer - User input passing the memory address to the variable

	fmt.Println("Enter number of tickets required: ")
	fmt.Scan(&userTickets) // Pointer - User input passing the memory address to the variable

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email) // Printing out user inputed information

}
