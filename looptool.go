/*
let we imagine we need to do the same thing again and again for
this situation we have one tool in our toolbox
loop toolkit is used when we want to perform same task again and again

it consists of for loop
range key word is used to loop the array,slce,map
*/
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("hi charan")
	}

	array := [5]int{1, 2, 3, 4, 5}
	for index, value := range array {
		fmt.Println(index, value)
	}
}
