package main

import "fmt"

func incrPtr(p *int) {
	//inc the actual value
	*p++
}

func swapVals(a, b *int) {
	*a, *b = *b, *a
}

func swapAddress(a, b *int) (*int, *int) {
	return b, a
}

func main() {
	var (
		val int
		ref int
		ptr *int
	)

	val = 0
	ptr = &val //pointer to val
	ref = *ptr //copy of val

	incrPtr(ptr)
	fmt.Println(val)
	fmt.Println(ref)

	a, b := 1, 2
	pa := &a
	pb := &b
	fmt.Printf("pa: %p — pb: %p\n", pa, pb)
	fmt.Printf("pa: %v — pb: %v\n", *pa, *pb)
	pa, pb = swapAddress(pa, pb)
	fmt.Printf("pa: %p — pb: %p\n", pa, pb)
	fmt.Printf("pa: %v — pb: %v\n", *pa, *pb)
	swapVals(pa, pb)
	fmt.Printf("pa: %p — pb: %p\n", pa, pb)
	fmt.Printf("pa: %v — pb: %v\n", *pa, *pb)
}
