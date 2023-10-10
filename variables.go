package main

import "fmt"

// let we take the first tool from the toolbox called variable
//we have lot of similar tools in our tool box for just storing the data but in different ways like const,maps,struct,array
//which are used to store the data temporarly  and these variables can be reused again and again there are like the small chips in machine

/*
these variables can store different type of data
int-integer
float -float32 and float64	
bool-boolean
string

Basic syntax
to declare the variable
var is the keyword use

var variable_name datatype or var varable_name (here lexer will take care about data type)

var name string


*/
func main() {
	var name string = "MR.C"
	//variable can be assigned in this two ways mostly we use := syntax
	name01 := "Mr.C"
	fmt.Println(name + " " + name01)
	//multiple variables can be assigned at a time by below syntax
	var a, h, j int = 1, 2, 3
	var b, c = 5, "hi bro"
	d, e := 2, "why bro"
	//constant are also the way to store the data but the values assigned fixed
	const f = 40
	

	fmt.Println(a, h, j, b, c, d, e, f)

}
