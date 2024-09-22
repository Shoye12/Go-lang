package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple","Potato","Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList,"Mango","Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	
	highScores := make([]int, 4)
	highScores[0] = 545
	highScores[1] = 445
	highScores[2] = 345
	highScores[3] = 245
	//highScores[4] = 045

	highScores = append(highScores, 123,564,823)

	//fmt.Println(highScores)

	//fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	//fmt.Println(highScores)

	var courses = []string{"ruby","javascript","swift","python","nodejs","reactjs"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) 
	fmt.Println(courses)

}