package main

import (
	"fmt"
)

var outA int 
var outB int = 2 
var outC = 3

func main() {
	// Variable declaration with an initial value
	var student1 string = "John"
	var student2 = "Jane"
	x := 2

	// Declaring variables without an initial value
	var a string   // Has a default value of "" Empty string
	var b int      // Has a default value of 0
	var c bool     // Has a default value of false

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	fmt.Println("----------------")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Assigning a value to a variable which was not initialized with a value
	fmt.Println("----------------")
	a = "James Bond"
	fmt.Println(a)

	/*
		OBS: It is not possible to declare a variable using ":=" without assigning a value to it

		Difference Between var and ":="

		var: 
			1. Can be used inside and outside of function
			2. Variable declaration and value assignment can be done separately 
		
		:= 
			1. Can be used inside function only
			2. Variable declaration and assignment must be done in the same line
	*/

	// Printing out side variables
	fmt.Println("")
	fmt.Println("------------ Out side variables")
	fmt.Println(outA)
	fmt.Println(outB)
	fmt.Println(outC)
}