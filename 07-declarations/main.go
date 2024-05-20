package main

import (
	"fmt"
	"path"
	"time"
)

//package scope declaration cant use inference syntax, must use var
var name = "timmy"

func main() {
	//normal
	var bool1 bool = false
	//short
	var bool2 = false
	//type inference
	bool3 := false
	name := "timmy"
	name2 := name
	str, i := "string", 1
	//in redeclaration, need to have atleast one new var
	str, newvar := "string", 2

	//swapping vars
	firstString, secondString := "first", "second"
	firstString, secondString = secondString, firstString

	_,_,_,_ = bool1, bool2,	bool3, name2
	
	fmt.Println(
		"str:", str,
		"i:", i,
		"newvar:", newvar,
	)

	fmt.Println(
		"firstString:", firstString,
		"secondString:", secondString,
	)

	//using time
	var now time.Time
	now = time.Now()
	fmt.Println(now)

	var file string
	_, file = path.Split("css/main.css")
	fmt.Println(file)
}