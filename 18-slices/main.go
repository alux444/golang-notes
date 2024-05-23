package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func slices() {
	// equivalent of lists
	// new slice is equivalent to nil
	// slices cant be compared
	var slice []int
	slice = append(slice, 1)

	slice2 := []int{1, 2, 3}
	//when the reaches over capacity, the slice gets overwritten with new one, as cap + len are immutable
	//so there is a new backing array for the slice.
	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice, slice2)

	// appending slices
	slice2 = append(slice2, slice...)
	fmt.Println(slice2)
	fmt.Println(slice2[:5])
}

func sliceSizes() {
	data := []string{"slices", "are", "awesome", "period", "!!" /* #5 */}
	fmt.Println(data)                                     // #1
	fmt.Printf("main's data slice's header: %p\n", &data) // #3
	array := [...]string{"slices", "are", "awesome", "period", "!!" /* #5 */}

	// array's size depends on its elements
	fmt.Printf("array's size: %d bytes.\n", unsafe.Sizeof(array))
	// slice's size is always fixed: 24 bytes (on a 64-bit system) — slice value = slice header
	// slice header is the pointer, size, capacity
	fmt.Printf("slice's size: %d bytes.\n", unsafe.Sizeof(data))
}

func underlying() {
	//slice has an underlying array from its state before
	slice := []int{1, 2, 3, 4, 5, 6}
	//slice is now [2,3] - it moves the FIRST reference of the slice to index 1 of array
	//the underlying array is [1,2,3,4,5,6], but slice start idx = 1
	slice = slice[1:3]
	fmt.Println(slice)
	//the subslice is just {1}, th
	slice = append(slice[:1], slice[:4]...)
	fmt.Println(slice)
	//underlying array is now [2,3,4,5,6] - keeps the overflow removed

	a := []int{1, 2, 3} //len = 3, cap = 3
	b := append(a[:1], 10)
	// a[:1] - len = 1, cap = 3
	// len is immutable
	// we update the next value to be '10'
	// from the underlying array : [1,2,3] -> [1,10,3], as it is shared between the slices
	// i.e both slices are pointing to ONE shared array
	fmt.Println(b)
	fmt.Println(a)

	arr := [...]int{9, 7, 5, 3, 1}
	nums := arr[2:]
	nums2 := nums[1:]

	arr[2]++
	nums[1] -= arr[4] - 4
	nums2[1] += 5

	//changes to nums, changes the backing array - shared array result, just with indices
	fmt.Println(nums)
}

func growth() {
	ages, oldCap := []int{1}, 1.

	for len(ages) < 5e5 {
		ages = append(ages, 1)

		c := float64(cap(ages))
		if c != oldCap {
			fmt.Printf("len:%-10d cap:%-10g growth:%.2f\n",
				len(ages), c, c/oldCap)
		}
		oldCap = c
	}
}

func fullSlice() {
	nums := []int{1, 3, 2, 4}
	//last val sets capacity - can set to smaller capacity.
	odds := nums[:2:2]
	evens := nums[:2:3]

	fmt.Println(nums)
	fmt.Println(odds)
	fmt.Println(evens)
	c := float64(cap(odds))
	fmt.Println(c)
	c = float64(cap(evens))
	fmt.Println(c)
}

func makes() {
	tasks := []string{"jump", "run", "swim", "fly"}

	//make a slice with type / length / capacity
	upTasks := make([]string, 0, len(tasks))

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		fmt.Println(upTasks)
	}
}

func copies() {
	evens := []int{2, 4, 6}
	odds := []int{3, 5, 7, 9, 11, 13, 15}

	fmt.Println("evens [before]:", evens)
	fmt.Println("odds  [before]:", odds)

	//copy will copy all elements from the source slice to the destination slice
	// N := copy(evens, odds)
	N := copy(odds, evens)
	fmt.Printf("%d element(s) are copied.\n", N)

	fmt.Println("evens [after]:", evens)
	fmt.Println("odds  [after]:", odds)
}

func twoDimensionalSlices() {
	slice := make([][]int, 2)
	subslice1 := []int{1, 2, 3}
	subslice2 := []int{4, 5, 6}
	subslice3 := []int{7, 8}
	slice = append(slice, subslice1, subslice2, subslice3)
	slice[0] = []int{1}
	fmt.Println(slice)
}

func main() {
	slices()
	sliceSizes()
	underlying()
	growth()
	fullSlice()
	makes()
	copies()
	twoDimensionalSlices()
}
