package mypackage

import "runtime"

//not exported
var one = 1
//exported
var Two = 2

//not exported
func test() string {
	return "test"
}

// GetVersion returns the Go version.
//capital letter - gets exported
func GetVersion() string {
	return runtime.Version()
}