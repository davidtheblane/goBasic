package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {
	
	greetUsers()

	for {
			firstName, lastName, email, userTickets := getUserInput()

			isValidaName, isValidEmail, isValidTicketNumber :=  validateUserInput(firstName, lastName, email, userTickets)

			if isValidaName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

				if remainingTickets == 0 {
					//end program
					fmt.Printf("Our conference is booked out. Come back next year.")
					break
				}
		} else {
			if !isValidaName{
				fmt.Printf("First name is invalid.\n")
			} 
			if !isValidEmail {
				fmt.Printf("Email is invalid.\n")
			}
			if !isValidTicketNumber {
				fmt.Print("number of tickets is invalid.\n")
			}
		}

	}
} 

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get yout tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
				for _, booking := range bookings {
					var names = strings.Fields(booking)
					firstNames = append(firstNames, names[0])
				}
				return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
			isValidaName := len(firstName) >= 2 && len(lastName) >= 2
			isValidEmail := strings.Contains(email, "@")
			isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
			return isValidaName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
			var firstName string
			var lastName string
			var email string
			var userTickets uint

			//ask user for their name
			fmt.Println("Enter your first name: ")
			fmt.Scan(&firstName)

			fmt.Println("Enter your last name: ")
			fmt.Scan(&lastName)

			fmt.Println("Enter your email: ")
			fmt.Scan(&email)

			fmt.Println("How much tickets you want?: ")
			fmt.Scan(&userTickets)
			return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName + " " + lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}