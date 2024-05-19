// Package main is a package
// go doc
package main

import (
	"fmt"
	"runtime"
)

var num1 = 10
var num2 = 5

func main() {
	if (num1 > num2) {
		fmt.Println("num1 is greater than num2")
	}
	fmt.Println("semicolon"); fmt.Println("test")
	fmt.Println(runtime.NumCPU())
}