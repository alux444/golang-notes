package main

import "fmt"

func main() {
	speed := 100
	force := 2.5

	//destructive conversion
	speed = speed * int(force)
	fmt.Println(speed)

	//non-destructive conversion
	speed = int(float64(speed) * force)
	fmt.Println(speed)

	//65 -> "A"
	a := 65 // 65 is A
	color := string(rune(a))
	fmt.Println(color)
	//bytes -> "hi"
	fmt.Println(string([]byte{104, 105}))

}