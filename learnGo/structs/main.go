package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

type Student struct {
	Name string
	Age  int
}
type Student2 struct {
	Age  int
	Name string
}

func main() {
	var a Student
	a.Name = "Oliver"
	fmt.Println(a)
	// 1 new keyword to create a new instance of a struct

	var pa *Student
	pa = new(Student)
	pa.Name = "Catherine"
	fmt.Println(pa)
	fmt.Printf("%T\n", pa)
	fmt.Printf("%#v\n", pa)
	fmt.Printf("%x\n", pa)
	fmt.Printf("%p\n", pa)
	// print 20 star
	fmt.Println(strings.Repeat("*", 20))
	b := Student{
		Name: "Oliver",
		Age:  21,
	}
	fmt.Println(b)
	fmt.Println(strings.Repeat("*", 20))
	// check alignment memeory
	fmt.Println(unsafe.Alignof(b))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Alignof(Student2{}))
	fmt.Println(unsafe.Sizeof(Student2{}))
	fmt.Println(unsafe.Alignof(Student{}))
	fmt.Println(unsafe.Sizeof(Student{}))
	i := 12
	fmt.Println("Integer alignment is", unsafe.Alignof(i))
	fmt.Println("the size of integer is", unsafe.Sizeof(i))
	s := "hello"
	fmt.Println("String alignment is", unsafe.Alignof(s))
	fmt.Println("the size of string is", unsafe.Sizeof(s))
	fmt.Println(strings.Repeat("*", 20))
	// structs are passed by value example
	fmt.Printf("struct are passed by value\n")
	fmt.Printf("Before b: %#v\n", b)
	c := b // c is a copy of b
	c.Name = "Catherine"
	fmt.Println("After b", b)
	fmt.Println("After c", c)
	fmt.Println(strings.Repeat("*", 20))
	d1 := Student{
		Name: "Oliver",
		Age:  21,
	}
	d2 := Student{
		Name: "Oliver",
		Age:  21,
	}
	fmt.Println(d1 == d2)

	fmt.Println(strings.Repeat("*", 20))
	// structs are passed by reference example
	fmt.Printf("struct are passed by reference\n")
	fmt.Printf("Before d1: %#v\n", d1)
	fmt.Printf("Before d2: %#v\n", d2)
	changeName(&d1)
	fmt.Printf("After d1: %#v\n", d1)
	fmt.Printf("After d2: %#v\n", d2)
	fmt.Println(strings.Repeat("*", 20))
	// structs are passed by value example
	fmt.Printf("struct are passed by value\n")
	fmt.Printf("Before d1: %#v\n", d1)
	fmt.Printf("Before d2: %#v\n", d2)
	changeName2(d1)
	fmt.Printf("After d1: %#v\n", d1)
	fmt.Printf("After d2: %#v\n", d2)
	fmt.Println(strings.Repeat("*", 20))
	inspectString("hello, 你好")

}

func changeName(s *Student) {
	s.Name = "Catherine"
}
func changeName2(s Student) {
	s.Name = "Oliver"
}

func inspectString(s string) {
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%x\n", s)

	// UTFMax is 4 -- up to 4 bytes per encoded rune
	var buf [utf8.UTFMax]byte

	// Iterate over the string
	for i, r := range s {
		// Get the bytes of the rune
		n := utf8.EncodeRune(buf[:], r)
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:n])
	}
	fmt.Println()
	for i, r := range s {
		// capture the bytes of the rune
		rl := utf8.RuneLen(r)

		si := i + rl

		copy(buf[:], s[i:si])

		// Display the details
		fmt.Printf("%2d: %q; codepoint: %#6x; encoded bytes: %#v\n", i, r, r, buf[:rl])
	}
}
