package main

import "fmt"

/*
let us imagine we have to take the different ways to go based on different conditions 
for example we have need to go right to school left to play ground based on mood 
we have to go in different ways

in this situation we have a wonderful toolpack in  our tool box called conditionaltoolpack

now we attached the storage chips with the operator tool packs.
let we take a look for another toolpack from our tool box ,
it is something special as it helps to perform different actions
based on different situations and conditions
these tool pack consists of tools
if else,if elseif else ,switch
it will use while taking the input from the user

*/

func main() {
	a := 18
	if a >= 18 {
		fmt.Println("you are allowed to visit")
	} else {
		fmt.Println("you are not allowed")
	}

	switch a {
	case 10:
		fmt.Println("you are not allowed")
	case 18:
		fmt.Println("you are allowed")
	default:
		fmt.Print("enter the input correctly")
	}

}
