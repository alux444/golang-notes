package main

import (
	"fmt"

	"github.com/alux444/golang-notes/05-library-package/mypackage"
)

func main() {
	fmt.Println(mypackage.GetVersion())
}