package main

import (
	"fmt"
	"runtime"
)

func main() {
	// NumCPU return the number of logical CPUs in the local machine.
	// It returns the number of cores if the operating system is Linux.
	// It returns the number of physical processors if the operating system is Windows.
	// It returns the number of logical processors if the operating system is Darwin.
	// It returns the number of physical processors if the operating system is FreeBSD.

	fmt.Println(runtime.NumCPU())
	println(runtime.GOOS)
}
