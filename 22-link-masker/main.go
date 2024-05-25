package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme somethin' to mask!")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text = args[0]
		size = len(text)
		buf  = []byte(text)
		in   bool
	)

	for i := 0; i < size; i++ {

		//if the length could fit, and the current index is equal to the pattern
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			i += nlink
		}

		//get char
		c := text[i]

		switch c {
		case ' ', '\n', '\t':
			in = false
		}

		if in {
			buf[i] = mask
		}
	}

	// print out the buffer as text (string)
	fmt.Println(string(buf))
}
