/*
Go is a new gen programming language designed for modren era
--compiled
-- minimalistic approach
--contains libraries which literally gives us the full controll from os

it is like a tool box for creating the cool machines in the digital world
let we just open up the tool box and familiarize with the tools and packages(bundled up tool for specific use)
we need to fix up the screws and the pipes inside the machine to get the desired output


let we know some terms
module- packages are grouped into module
package - code is grouped into package

Let dive with this simple example

steps to follow
create a module by creating a folder
enable dependency tracking by using(go mod init modulename)
declare the package 
*/

package main

//every code in go belongs to a package
//main package is the entry point of every module where we write the entry point of the machine
//when we start the machine main function is first function which calls up.

import "fmt"

//fmt is the prebuild tool in the toolbox with specific functionalities like print to screen ...etc
//which we taken from the tool box and directly fitted into the machine

func main() {
	//function is a block of code designed to perform a specific task when ever it called
	//it is like a solderd board with wires capacitors transistors ..etc
	//let we discuss
	fmt.Printf("Hello Universe")
}
