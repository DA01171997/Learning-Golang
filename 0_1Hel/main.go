package main

import "fmt"

func main() {
	var name = "YO"
	var name1 = 23
	var name2 = true
	const yo = "hi"
	name3 := 21
	size := 1.3
	name, age, weight := "duy", 21, 2.3
	fmt.Printf("%T,%T,%T\n", name, age, weight)
	fmt.Println(name)
	fmt.Println("hello world", name)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", name1)
	fmt.Printf("%T\n", name2)
	fmt.Printf("%T\n", name3)
	fmt.Printf("%T\n", size)
	fmt.Printf("%T\n", yo)
}
