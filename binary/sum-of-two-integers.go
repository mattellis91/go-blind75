package main

//Given two integers a and b, return the sum of the two integers without using the operators + and -.

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Print(getSum(a, b))
}

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	return getSum(a^b, (a&b)<<1)
}