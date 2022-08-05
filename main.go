// Go CLI learning file
// To run - go run main.go
// TO run all of a file, run go run .
// Comment with lines to understand application.
//FMT - Println (Prints lines of code), Printf (Prints line with variables declared after), %v (Variable), %T (Prints data type.)

package main // Declaring package, called main which is the common one.

import (
	"booking-app/helper" // Importing local application package, helper
	"fmt"                //Importing packages using a list of supplied ones. fmt is common package.
	"sync"               // Importing sync
	"time"               // Importing time
)

const conferenceTickets int = 50     // Declaring a constant (value dosen't change)
var conferenceName = "Go Conference" // Declaring a variable with alternative syntax
var remainingTickets uint = 50       // Declaring a variable
var bookings = make([]UserData, 0)   //Intialising a list of maps

type UserData struct { // Creating a struct with multiple data
	firstName       string // Data with variable
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{} // Creating a variable called wg with sync function to wait to enable the application to support multi threaded workloads or it would finish in one

func main() { // Declaring a function called main

	greetUsers() //Using GreatUsers function

	//Function Logic
	//For loop to continue running the applicatioj
	//User input validation
	firstName, lastName, email, userTickets := getUserInput()                                                                             // Local function with variables
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets) //Imported package function

	//Remaining ticket logic

	if isValidName && isValidEmail && isValidTicketNumber { //Logic for if a user tries to book more tickets than available

		bookTicket(userTickets, firstName, lastName, email)    // Local function with variables
		wg.Add(1)                                              // Adding a waiting group (1) to enable more user input while waiting for email proccessing
		go sendTicket(userTickets, firstName, lastName, email) //go enables go concureency multi threaded

		firstNames := getFirstNames() //Creating a variable with a local function
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		//Checking if tickets are sold out and weather to continue
		if remainingTickets == 0 {
			//End Program
			fmt.Println("Our conference is booked out. Come back next year!")
		}
	} else { //Letting user which input failed.
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

	wg.Wait()
}

//External function
func greetUsers() {

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still avalable to purchase. \n", conferenceTickets, remainingTickets) // Using FMT format package to add a placeholder for the variable and declare it later
	fmt.Println("Get your tickets here to attend")                                                                             // Using Println to print text
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // Gives the index and element value
		firstNames = append(firstNames, booking.firstName) //Slice
	}

	return firstNames

}

func getUserInput() (string, string, string, uint) {
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

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets // logic to calculate remaning tickets

	//User data map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData) // booking variable slice which uses append() which is a built in go function with the array, and variables wanted to pass in.
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email) // Printing out user inputed information
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)                                                                // Printing remaining ticktes using memory logic
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // Adding a time delay
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")
	wg.Done() //Wait Group done command
}
