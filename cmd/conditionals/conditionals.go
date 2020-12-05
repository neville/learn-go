package main

import "fmt"

func main() {
	// if-else-if statements
	// else statement should always start in the same line after the closing if
	number := 10

	if number%2 == 0 {
		fmt.Println("number", number, "is even")
	} else {
		fmt.Println("number", number, "is odd")
	}

	// variant with an optional statement
	// scope of 'success' is restricted only to the if-else block
	if success := false; success == true {
		fmt.Println("success is", success)
	} else {
		fmt.Println("success is", success)
	}

	// switch and case
	mode := 1
	switch mode {
	case 1:
		{
			fmt.Println("Mode is ON")
		}
	case 0:
		{
			fmt.Println("Mode is OFF")
		}
	}

	// variant with an optional statement
	// without brackets for cases and a default case
	// scope of 'dayOfTheWeek' is restricted only to the if-else block
	switch dayOfTheWeek := 7; dayOfTheWeek {
	case 1:
		fmt.Println("It's Monday!")
	case 2:
		fmt.Println("It's Tuesday!")
	case 3:
		fmt.Println("It's Wednesday!")
	case 4:
		fmt.Println("It's Thursday!")
	case 5:
		fmt.Println("It's Friday!")
	case 6:
		fmt.Println("It's Saturday!")
	case 7:
		fmt.Println("It's Sunday!")
	default:
		fmt.Println("Which world are you in!!!")
	}

	// variant with multiple expression in a case

	// switch without an expression

	// fallthrough case
}
