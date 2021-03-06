package main

import "fmt"

func main() {
	// Declaring a variable
	var name string

	name = "Neville"
	fmt.Println("My name is", name)

	name = "Luke"
	fmt.Println("Now my name is", name)

	// Declaring a variable with an initial value
	var luckyNumber int = 29
	fmt.Println("My lucky number is", luckyNumber)

	// Type inference
	var degree = 37.3
	fmt.Println("The degree is", degree)

	// Multi variable declarations
	var length, breadth int = 17, 19
	fmt.Println("Length and breadth is", length, "and", breadth)

	var area, perimeter = 15.45, 13.99
	fmt.Println("The area and perimeter is", area, "and", perimeter)

	// Default values of variables if there is no initial value specified
	var age int
	var temperature float32
	var lastName string
	fmt.Println("Values in variables age, temperature and lastName are ", age, " ", temperature, " ", lastName)

	// You can also declare variables in the following syntax
	var (
		country          = "India"
		pincode          = 400053
		distance float32 = 414.55
	)
	fmt.Println("country is", country, "and pincode is", pincode, "and distance is", distance)

	// Short hand declaration
	// Should always have inital values
	item := "beer"
	fmt.Println("I have a", item)

	// This syntax works only when atleast one variable is newly declared
	item, amount := "cocktail", 90
	fmt.Println("This", item, "costs around", amount, "rupees.")
}
