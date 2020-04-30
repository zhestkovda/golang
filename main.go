package main

import (
	"fmt"
	"strconv"
)

const pi float64 = 3.1415

func main() {

	var hello string
	k := 5
	var a int = 4
	var b int = 3
	if a < b {
		fmt.Println("a меньше b")
	} else {
		fmt.Println("a больше b")
	}

	var c bool = a == b
	_ = c
	//fmt.Println(c)

	hello = "привет"
	fmt.Println(hello + " " + strconv.Itoa(k))

	var numbers [5]int = [5]int{1, 2}
	fmt.Println(numbers) // [1 2 0 0 0]
}
