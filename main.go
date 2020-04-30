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
	fmt.Println(c)

	hello = "привет"
	fmt.Println(hello + " " + strconv.Itoa(k))

	var numbers [5]int = [5]int{1, 2}
	fmt.Println(numbers) // [1 2 0 0 0]

	aq := 5
	switch aq {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	case 6, 5, 4:
		fmt.Println("a = 6 или 5 или 4, но это не точно")
	default:
		fmt.Println("значение переменной a не определено")
	}

	
}
