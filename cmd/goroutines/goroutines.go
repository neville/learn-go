package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello there...")
}

/*
 * The main function always runs in its own goroutine called the main goroutine
 * It should always be running for any other goroutines to run
 * Once the main goroutine terminates, the program will be termianted and no other goroutines will run
 */
func main() {
	/*
	* When a new goroutine is started, the call returns immediately
	* It does not wait for the goroutine to finish executing
	* Any values returned from the goroutine are ignored
	 */
	go sayHello()

	go printNumbers()

	go printCharacters()

	/*
	* Sleep is a hack used here to make the main goroutine wait for other gorutines to finish their execution
	* If we didnt use it, then the control would return immediately to 'main' without waiting for the other goroutine to finish executing
	* 'main' then finishes executing without giving the other gorutines any time to finish running
	 */
	//time.Sleep(3 * time.Second)

	fmt.Println("main terminating...")
}

func printNumbers() {
	for i := 0; i <= 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func printCharacters() {
	for c := 'a'; c <= 'e'; c++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c\n", c)
	}
}
