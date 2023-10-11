/*
let we take another special tool from our tool box which will help to specify our own data type
like int for integer,float32 for float number we can specify our own datatype with this tool
this tool is called struct

to define the struct
type keyword is used

type student struct{
	name string
	rollno int
	grade string
}

var s1=student{"mrc",27,"a1"}
s1:=student{name:"mrc",rollno:27,grade:"a1"}

we can embed the functiontool in struct it is called method now

func(s student) password() string{
	return name+grade
}

*/

package main

import "fmt"

type Student struct {
	name   string
	rollno int
	grade  string
}

func (s Student) password() string {
	return s.name + s.grade
}
func main() {

	s1 := Student{name: "mrc", rollno: 27, grade: "a1"}
	fmt.Println(s1)
	fmt.Println(s1.password())
}
