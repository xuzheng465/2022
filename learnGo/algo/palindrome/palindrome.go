package main

import (
	"fmt"
	"strconv"
)

func IsPalindrome(x int) bool {
	return false
}

func main() {
	x := 121
	sx := strconv.Itoa(x)

	for i := 0; i < len(sx); i++ {
		fmt.Println(sx[i])
	}
}
