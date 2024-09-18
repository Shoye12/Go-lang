package main

import (
	"fmt"
)
const LoginToken string = "skmoevmeo"  // public
func main() {
	var username string = "Shoye"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var smallFloat float64 = 255.54569654456441565322
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
} 