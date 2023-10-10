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

mostly we don't use this chip array in building our machine
we use another chip from our toolbox called slices
slices are the arrays but more flexible and powerful

syntax is similar to arrays but no size is specified

var slice_name []datatype

we can use make keyword to create the slice

slice_name:=make([]datatype,length,capacity)

assigning

var slice_name=[]int{1,2,3}
slice_name:=[]int{1,2,3}
we can peform some functions on slices like
len()
cap()
append()

we can even transfer the full or partial data from array to slices by using below syntax
slice_name[startindex:endindex]


*/

func main() {
	var coin_types = [3]string{"one ruppe", "two ruppe", "three ruppe"}
	fmt.Println(coin_types[1])
	titles := make([]string, 3)
	titles = append(titles, "ghost", "demon", "rock")
	fmt.Println(titles)
	fmt.Println(coin_types[0:1])

}
