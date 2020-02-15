package main

import "fmt"

func main() {
	//Arrays
	var fruitArr [2]string

	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	//Declare and assign
	fruitArrShort := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr, fruitArrShort)
	fmt.Println(fruitArr[1])

	fruitArrSlice := []string{"Apple", "Orange", "Grapes"}

	fmt.Println(fruitArrSlice)
	fmt.Println(len(fruitArrSlice))
	fmt.Println(fruitArrSlice[1:2])

}
