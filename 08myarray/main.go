package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList[5]string

	fruitList[0] = "Banana"
	fruitList[1] = "Apple"
	fruitList[2] = "Grapes"

	fmt.Println("Fruitlist is:",fruitList)
	fmt.Println("Fruitlist is:",len(fruitList))
	var vegList = []string{"capsicum","potato","brinjal"}
	fmt.Println("vegList is:",(vegList))
	fmt.Println("vegList is:",len(vegList))

}