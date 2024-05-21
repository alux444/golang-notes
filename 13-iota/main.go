package main

import "fmt"

func main() {
	//iota starts at 0, increments by 1
	const (
		monday = -5 + iota
		tuesday
		_ //skip
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
}