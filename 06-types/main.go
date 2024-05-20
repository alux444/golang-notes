package main

import "fmt"

func main() {
	var integer int = 1
	var boolean bool = false
	var str string = "hello"
	var hexa int = 0x1A

	var int16 int16 = 1
	var int32 int32 = 1
	var int64 int64 = 1
	var float32 float32 = 1.1
	var float64 float64 = 1.1
	var complex64 complex64 = 1 + 1i
	var complex128 complex128 = 1 + 1i
	
	// string and also numeric
	var rune rune = 'a'
	var byte byte = 'a'

	//multiple declaration
	var (
		one int = 1
		two string = "2"
	)

	//keeping unused variable
	var unused bool
	_ = unused

	fmt.Println(integer)
	fmt.Println(boolean)
	fmt.Println(str)
	fmt.Println(hexa)
	fmt.Println(int16)
	fmt.Println(int32)
	fmt.Println(int64)
	fmt.Println(float32)
	fmt.Println(float64)
	fmt.Println(complex64)
	fmt.Println(complex128)
	fmt.Println(rune)
	fmt.Println(byte)

	//print type and val
	fmt.Printf("%T %v %T %v\n", one, one, two, two)
}