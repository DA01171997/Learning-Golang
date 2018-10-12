package main

import (
	"fmt"
)

func meeting(name string) string {
	return name + "hi"
}
func getSum(one, two int) int {
	return one + two
}

func main() {

	fmt.Println(2, meeting("hi"))
	fmt.Println(2, getSum(23, 23))
}
