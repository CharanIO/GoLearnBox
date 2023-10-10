package main

import "fmt"

/*
In this Storage chips section ,
we will take a look with all the chips that are used to storage, form our tool box

the are the tools we will mess up
->array
->slice
->maps

arrays are the variables too
but they are used to store multiple values of same datatype in a comman name but fixed length

syntax

var array_name [length]int
array_name:=[length]int{values}
while assigining

var array_name=[length]int{values}

for acessing the values the with the index values starts from 0

array_name[0]....


*/

func main() {
	var coin_types = [3]string{"one ruppe", "two ruppe", "three ruppe"}
	fmt.Println(coin_types[1])
}
