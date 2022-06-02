package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"a", "b", "c"}

	i := 1
	// remove the element at index i from a
	a = append(a[:i], a[i+1:]...)
	fmt.Printf("length of a is %d\n", len(a))
	fmt.Printf("capacity of a is %d\n", cap(a))
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5}
	j := 2
	// remove the element at index j from b
	copy(b[j:], b[j+1:])
	//b = b[:len(b)-1]
	fmt.Println(b)

	// print 20 lines of "-"
	fmt.Println(strings.Repeat("-", 20))

	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Orange"
	slice1[2] = "Banana"
	slice1[3] = "Grape"
	slice1[4] = "Pineapple"
	inspectSlice(slice1)
	fmt.Println(strings.Repeat("-", 20))
	// Take a slice of slice1. We want just indexes 2 and 3.
	// Parameters are start, end (exclusive).
	// Parameters are [start_index: start_index + length]
	slice2 := slice1[2 : 2+2]
	inspectSlice(slice2)
	fmt.Println(strings.Repeat("-", 20))

	slice3 := slice1[2:4:4]
	inspectSlice(slice3)

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1[1] = 100
	fmt.Println(s1)
	changeSlice(s1)
	fmt.Println(s1)

}
func changeSlice(slice []int) {
	slice[2] = 1000
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %patterns.go\n", i, &slice[i], s)
	}
}
