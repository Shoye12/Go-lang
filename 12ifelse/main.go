package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount := 23
	var result string
	if loginCount < 10{
		result = "Regular user"
	} else if loginCount >10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2==0{
		fmt.Println("Number is even")
		}else {
			fmt.Println("number is odd")
		}
		if num := 3{

		}
	// if err != nil {
		
	// }
}