package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	var num = rand.Intn(10) + 1
	fmt.Printf("%d\n", num)
}
