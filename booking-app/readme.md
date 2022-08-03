### install GO (https://www.cyberciti.biz/faq/how-to-install-gol-ang-on-ubuntu-linux/)
sudo apt install golang

#### Initialize a project in GO
go mod init application-name


### Go works with imported packages
import "fmt"
like "fmt" where you can run fmt.Print('Hello World')

# package main is the main package and always need to be declared

# func main() is the main function and serves as entrypoint, she has to be declared only one time in the code.

# VARIABLES
declare a variable in go
| variableType variable name = variableValue |
EXAMPLE: var name = "João

var could have his value altered, but const not
const cpf = '9999999999'

# Passing variables in strings
using Println => fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
using printf =>  fmt.Printf("Welcome to %v booking application \n", conferenceName)Variables

# DATA TYPES
Initialize a Variable
var varName varDataType

Strings
Booleans
Integers

works with Printf only
print variables value = %v
print variables data type = %T

# Compared to Java
Go       Java
int8  => byte
int16 => short
int32 => int
int64 => long

uInt = whole positive numbers 

Float 
not whole numbers

# "Syntactic Sugar" way of declare a variable and assign
# Can't use with constants
instead of 
var conferenceName = "Go Conference"
use
conferenceName := "Go Conference"

# What is a pointer? "&"
pointer is a variable that points to a memory value of another variable tha references the actual value.
in go we call as

fmt.Scan(&userName) 
fmt.Scan assign the users input value
"&" is a pointer of var userName

# Arrays
Arrays have fixed Size, so we have to pass the value size
We have to declare the variables types
Only store the same data type
# sintax
var array_name[size]variable_type{"array itens"}
# We could declare an empty array like this
var array_name [size]type

# Slices,
is an abstraction of an Array
Automatically expands when new elements are added
# sintax
var slice []variable_type
var slice = []variable_type{}
slice := []variable_type{}
to send an item to slice we use append()
# Example:
bookings = append(array, content)

# See a length of an element
len(variable)

# LOOPS IN GO

in GO we only have FOR loop
sintax
    firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
		}
# range 
iterates over elements for difeenrent data structures
For arrays and slices, range provides the index and value for each element
# split
strings.Fields(var)
splits the string with white space as separator
returns a slice with the split elements
# underscore
when we need to declare a var but not use it.

# IF else statements
sintax
if condition {
  declaration
} else {

}

# Validate User Input
isValidaName := len(firstName) >= 2 && len(lastName) >= 2
isValidEmail := strings.Contains(email, "@")
isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

# function
in a function we have to specify the type of value we are passing and returning
example:
											\/input value			\/output value
func printFirstNames(bookings []string) []string{
	firstNames := []string{}
				for _, booking := range bookings {
					var names = strings.Fields(booking)
					firstNames = append(firstNames, names[0])
				}
				return firstNames
}

# Go return multiple values in a function
example:
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
			isValidaName := len(firstName) >= 2 && len(lastName) >= 2
			isValidEmail := strings.Contains(email, "@")
			isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
			return isValidaName, isValidEmail, isValidTicketNumber
}

# Package level variables
variables defined at the top, outside all functions

# GO PACKAGES
We could run function in other files in go just passing their dependencies and "package main"

# exporting a function in go
just capitalize the first letter
serviceFunction turns into ServiceFunction

# maps
all keys and values have the same type
# sintax to create a empty map
var map = make(map[string]string)

#sintax to create a list of empty maps
var listOfMaps = make([]map[string]string, 0)

# convert userTickets in string
	strconv.FormatUint(uint64(userTickets), 10)

# structs
create a struct
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

#concurrency


PAREI EM 3:10:36