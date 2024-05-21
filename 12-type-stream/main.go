package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"
)

type (
	// Gram underlying type is int
	Gram int
	// Kilogram underlying type is int
	Kilogram Gram
	// Ton underlying type is int
	Ton Kilogram
)
// Gram and weights.Gram are different types

// You cannot do this
// salt = weights.Gram(100)

// But, you can convert it back to Gram
// Since, their underlying type is int

func bitsAndBytes() {
	//%b - bits
	fmt.Printf("%b\n",0)
	fmt.Printf("%b\n",1)

	//2 bits with leading 0s
	fmt.Printf("%02b\n",0)
	fmt.Printf("%02b\n",1)

	//8 leading
	fmt.Printf("%08b\n",0)
	fmt.Printf("%08b\n",1)

	//bits to int
	i, _ := strconv.ParseInt("00010110", 2, 64)
	fmt.Println(i)

	//bytes - int8 (0-255)
	fmt.Printf("%08b = %d\n", 255, 255)
}

func mathFuncs() {
	fmt.Println("int8:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, math.MaxInt64)

	//unsigned
	fmt.Println("uint8:", 0, math.MaxUint8)
	fmt.Println("uint16:", 0, math.MaxUint16)
	fmt.Println("uint32:", 0, math.MaxUint32)
	fmt.Println("uint64:", 0, uint64(math.MaxUint64))
	//max unsigned64 can only be used in constant expressions

	fmt.Println("float32:", math.SmallestNonzeroFloat32,
		math.MaxFloat32)
	fmt.Println("float64:", math.SmallestNonzeroFloat64,
		math.MaxFloat64)

	// memory costs
	fmt.Println("int    :", unsafe.Sizeof(int(1)), "bytes")
	fmt.Println("int8   :", unsafe.Sizeof(int8(1)), "bytes")
	fmt.Println("int16  :", unsafe.Sizeof(int16(1)), "bytes")
	fmt.Println("int32  :", unsafe.Sizeof(int32(1)), "bytes")
	fmt.Println("int64  :", unsafe.Sizeof(int64(1)), "bytes")

	fmt.Println("uint   :", unsafe.Sizeof(uint(1)), "bytes")
	fmt.Println("uint8  :", unsafe.Sizeof(uint8(1)), "bytes")
	fmt.Println("uint16 :", unsafe.Sizeof(uint16(1)), "bytes")
	fmt.Println("uint32 :", unsafe.Sizeof(uint32(1)), "bytes")
	fmt.Println("uint64 :", unsafe.Sizeof(uint64(1)), "bytes")

	fmt.Println("float32:", unsafe.Sizeof(float32(1)), "bytes")
	fmt.Println("float64:", unsafe.Sizeof(float64(1)), "bytes")
}

func times() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Printf("%T\n", now)

	h, _ := time.ParseDuration("4h30m")
	fmt.Println(h)
	fmt.Printf("%T\n", h)
	fmt.Printf("%T\n", int64(h))
}

func main() {
	bitsAndBytes()
	mathFuncs()
	times()
}