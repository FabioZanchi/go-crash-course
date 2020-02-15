package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// use to read val from address/the memory
	fmt.Println(*b)
	fmt.Println(&a)
	fmt.Println(*&a)

	//Change val with Pointer
	*b = 10
	fmt.Println(a)

}
