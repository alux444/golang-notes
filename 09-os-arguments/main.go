package main

import (
	"fmt"
	"os"
)

func main() {
	//go run main.go arg1 arg2 arg3
	if len(os.Args[1:]) < 1 {
		fmt.Println("need at least one argument")
		os.Exit(1)
	}

	fmt.Println(len(os.Args[1:]))
	fmt.Println(os.Args[1:])
	fmt.Println("path:", os.Args[0])
}