package main

import "fmt"

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a int, b int, c int) int {
	return plus(a, b) + c
}
