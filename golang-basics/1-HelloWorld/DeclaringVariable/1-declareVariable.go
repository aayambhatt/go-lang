package main

import "fmt"

func main(){
	var myName string
	myName = "Aayam Bhatt"
	fmt.Println("The name of that backend engineer: ", myName)

	// short way to declare variable, and combining bool, int, string in one sentence
	isAdmin := true
	characterAge := 45

	fmt.Printf("USER_NAME: %s \nADMIN_STATUS: %t \nUSER'S AGE: %d\n", myName, isAdmin, characterAge)
}