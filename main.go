package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello world", validate("cassie", 5))
	fmt.Println("time", time.Now())
}

func swap(x, y string) (string, string) {
	return y, x
}

func validate(input string, number int) int {
	if input == "cassie" {
		return 4 * number
	}
	return 0 * number
}