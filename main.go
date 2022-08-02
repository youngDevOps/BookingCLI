// Go CLI learning file
// To run - go run main.go
// Comment with lines to understand application.

package main // Declaring package, called main which is the common one.

import "fmt" //Importing packages using a list of supplied ones. fmt is common package.

func main() { // Declaring a function called main
	var conferenceName = "Go Conference" // Declaring a variable
	const conferenceTickets = 50         // Declaring a constant (value dosen't change)
	var remainingTickets = 50            // Declaring a variable

	fmt.Println("Welcome to", conferenceName, "booking application")                                                        // Using FMT Package to print a line with text and variable
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still avalable to purchase") // Using FMT Package to print a line with text and variable
	fmt.Println("Get your tickets here to attend")                                                                          // Using Println to print text

}
