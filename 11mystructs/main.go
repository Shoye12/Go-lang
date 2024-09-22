package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	//no inheritance in golang; No super no parent

	Shoye := User{"Shoye","shaikshoyeb30@gmail.com",true,23}
	fmt.Println(Shoye)
	fmt.Printf("Shoye details are:%+v\n",Shoye)
	fmt.Printf("Name is %v and email is %v",Shoye.Name,Shoye.Email)
}
	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int

	}

