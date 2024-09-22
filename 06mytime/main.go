package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Go-lang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 Monday"))
	createdDate := time.Date(2024,time.August,21,23,25,55,56,time.UTC)
	fmt.Println(createdDate)
	var buf bytes.Buffer
	fmt.Fprintln(&buf,createdDate.Format("02-01-2006 Monday"))
}