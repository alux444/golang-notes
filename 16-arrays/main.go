package main

import "fmt"

func arrays() {
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	var arr2 = [2]int{1, 2}
	arr3 := [3]int{1, 2}
	arr4 := [...]int{1, 2, 3, 4, 5}

	_, _, _, _ = arr, arr2, arr3, arr4
	//same type, same values are equal
	fmt.Println(arr == arr2)
}

func twoDimesionalArrays() {
	arr := [...][2]int{{1, 2}, {3, 4}}

	sum := 0
	for index, value := range arr {
		for index2, value2 := range value {
			sum += value2
			_ = index2
		}
		_ = index
	}
}

func keys() {
	//keyed array
	arr := [...]int{
		1: 10,
		5: 20,
		21,
		0: 2,
	}
	fmt.Println(arr)

	//using iota for keys
	const (
		A = 9 - iota
		B
		C
	)

	rates := [...]float64{
		A: 25.5,
		B: 120.5,
		C: 20,
	}

	fmt.Printf("%#v\n", rates)
}

func main() {
	arrays()
	twoDimesionalArrays()
	keys()
}
