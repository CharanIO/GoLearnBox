/*
we recently messed up with the tools for storage now
now we setup the the datastorage tools from the tool box
now we need to connect and controll the flow by fixing up the screws and pipes
we need to perform some operation on this storage chips

operations we have different tool packs in our tool box like

arithmetic
assignment
comparison
logical
bitwise

arithmetic tool pack used for performing mathematical operations on the data

	which is stored in storage chips

we have (+,-,*,%,++,--)components which we can use with the storage tools

comparision tool pack we have (<,>,<=,>=,<> or!=) for comparing values of variables

assignment tool pack we have (==,+=,-=,*=,%=,!=,^=) used for assigning values to variables

logical tool pack for logical operations we have (&&,||,!)

bitwise tool pack
*/
package main

import "fmt"

func main() {
	//here we take the small chips from the toolbox perform some operation from operation toolkit 
	//to store the values in it and we use the prebuild tool (package) fmt 
	//to print the value to screen
	a := 10
	b := 20
	b += 10
	c := a * b
	fmt.Println(c)
	y := a > 2 && b <= 20
	fmt.Print(y)

}
