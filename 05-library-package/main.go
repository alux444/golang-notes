package main

import (
	"fmt"

	mypackage "github.com/my-org/my-package/util"
)

func main() {
	fmt.Println(mypackage.GetVersion())
	fmt.Println(mypackage.Two)
}