package helper //Declaring a new package to split up the application

import "strings" //Impoting Strings application

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) { //Function with input values and then return paramaters
	isValidName := len(firstName) >= 2 && len(lastName) >= 2                  //checking length of string input
	isValidEmail := strings.Contains(email, "@")                              //checking if the string contains a @ character and returns in a boolean format
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets //Checking user tickets
	return isValidName, isValidEmail, isValidTicketNumber                     // Returning the values
}
