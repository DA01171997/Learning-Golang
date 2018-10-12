package main

import (
	"fmt"
	"math"

	"github.com/DA01171997/go_crash_course/02_package/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(6))
	test := "hihihi"
	fmt.Println(strutil.Reverse(test))
}
