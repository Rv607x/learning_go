package main

import "fmt"

func main(){
	//ints are initialized with zero if no value is assigned to them
	var roomNumber, floorNumber int
	fmt.Println("room number and floor number are:", roomNumber, floorNumber)
	
	//strings are initialized with zero but outputs an empty string if no value assigned
	var password string
	fmt.Println("password is:", password)
}