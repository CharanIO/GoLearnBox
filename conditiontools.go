package main

import "fmt"

/*

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
