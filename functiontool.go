/*
let we take out the another tool from our toolkit ,it looks like a PCB(printed circuit board)
on which we can attach other tools like storage tools,operation tools etc
it is called function
function is a block of code to perform specific task when ever it is called and
we can pass the values to it (parameters)

syntax
func function_name(var_name datatype){

}

calling the function
function_name(var_value)

function can return the values out
by using return in the function we can return values out
*/

package main

import "fmt"

func main() {
	value := sum(3, 4)
	fmt.Println(value)

}

func sum(x int, y int) int {
	return x + y
}

//recursive function is the type of function where function call itself untill the base condition is reaced
