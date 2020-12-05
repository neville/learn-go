package main

import "fmt"

/*
 * goroutines communicate with each other using channels
 */
func main() {
	/*
	* every channel has a type <T> associated with it
	* the zero value of a channel is nil
	* which is why we need to define it like we do for slices and maps
	 */
	var numbers chan int

	if numbers == nil {
		numbers = make(chan int)
		fmt.Printf("Numbers is of type %T\n", numbers)
	}

	// a channel can also be created using a shorthand notation
	characters := make(chan string)
	fmt.Printf("Characters is of type %T\n", characters)

	/*
	* Sending and receiving data from channels
	* Sending and receiving data to and from channels is blocking by default
	* no explicit locks are required
	 */
	go getANumber(numbers)
	go getALetter(characters)

	// Read data from a channel
	data := <-numbers
	fmt.Printf("Value in data is %d\n", data)

	letter := <-characters
	fmt.Printf("Value in letter is %s\n", letter)
}

func getANumber(done chan int) {
	fmt.Println("I'm passing the number 555")

	// Write data to a channel
	done <- 555
}

func getALetter(done chan string) {
	fmt.Println("I'm passing the letter N")

	// Write data to a channel
	done <- "N"
}
