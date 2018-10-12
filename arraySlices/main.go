package main

import "fmt"

func main() {
	//Array
	//var fruitArr [2]string
	//fruitArr[0] = "apple"
	//fruitArr[1] = fruitArr[0]

	//Declare and assign

	//fruitArr := [2]string{"apple", "apple"}

	//fmt.Println(fruitArr)

	fruitslice := []string{"Apple", "Orange", "Coconu"}

	fmt.Println(len(fruitslice))
	fmt.Println(fruitslice[0:2])
}
