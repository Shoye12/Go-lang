package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Welcome to maths in Golang")

	//var mynumberOne int = 2
	//var mynumberTwo float64 = 4

	//fmt.Println("The sum is:",mynumberOne+int(mynumberTwo)

	// random number
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//fmt.Println(r.Intn(5)+1)

	myRandomNum, _ := rand.Int(rand.Reader,big.NewInt(5))
	fmt.Println(myRandomNum) 


	
}