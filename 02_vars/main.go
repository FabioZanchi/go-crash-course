package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	var name string = "Fabio"
	var age int = 34
	var isCool = true
	isCool = false

	//Using const
	const isCoolConst = true

	fmt.Println(name, age, isCool, isCoolConst)
	fmt.Printf("%T\n", isCool)

	// Shorthand method used only inside functions
	shorthandName := "Fabio"
	size := 1.3

	shortName, shortAge := "Fabio", 32

	fmt.Println(shorthandName, size, shortName, shortAge)
}
