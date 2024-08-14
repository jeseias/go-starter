package main

import (
	"fmt"
)

var outA int 
var outB int = 2 
var outC = 3

func output() {
	var i,j = "hello", "world"

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	fmt.Print(i, "\n", j)

	fmt.Print(i, " ",j)
}

func constants() {
	const PI = 3.14

	fmt.Println(PI)
}

func moreVariables() {
	fmt.Println("")
	fmt.Println("----------------------")
	fmt.Println("")
	fmt.Println("More About variables")

	// Declaring multiple variables
	var m1, m2, m3, m4 int = 1, 2, 3, 4

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println(m4)

	// Multiple types on variables
	var x1, x2 = 5, "James"
	x3, x4 := 3, "Bond"

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
}

func firstVars() {
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

func main() {
	
	firstVars()
	moreVariables()
	constants()

	output()
}