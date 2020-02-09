package main

import "fmt"

func main() {
	// bool
	success := true
	completed := false
	fmt.Println("success is", success, "and completed is", completed)

	// signed integers
	// integer
	// size is either 32 or 64 bits and it depends on the underlying platform
	var length int = 567
	var temperature = -238
	fmt.Println("length is", length, "and temperature is", temperature)

	//  int8
	var breadth int8 = 127
	var width int8 = -128
	fmt.Println("breadth is", breadth, "and width is", width)

	// int16
	var height int16 = 32767
	var volume int16 = -32768
	fmt.Println("height is", height, "and volume is", volume)

	// int32
	var distance int32 = 2147483647
	var time int32 = -2147483648
	fmt.Println("distance is", distance, "and time is", time)

	// int64
	var speed int64 = 9223372036854775807
	var area int64 = -9223372036854775808
	fmt.Println("speed is", speed, "and area is", area)

	// unsigned integers
	// uint
	// size is either 32 or 64 bits and it depends on the underlying platform
	var centimeters uint = 968
	fmt.Println("Value of centimeters is", centimeters)

	// uint8
	var meters uint8 = 255
	fmt.Println("Value of meters is", meters)

	// unit16
	var kilometers uint16 = 65535
	fmt.Println("Value of kilometers is", kilometers)

	// uint32
	var milimeters uint32 = 4294967295
	fmt.Println("Value of milimeters is", milimeters)

	// uint64
	var inches uint64 = 18446744073709551615
	fmt.Println("Value of inches is", inches)

	// floating points

	// complex

	// byte

	// rune

	// string

	// type conversion
}
