package main

import "fmt"

func add1(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add1(42, 133))
	fmt.Println(add2(42, 133))
}
