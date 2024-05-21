package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	word := "hello"
	for i, v := range word {
		fmt.Println(i, string(v))
	}

	words := strings.Fields(corpus)
	query := os.Args[1:]

	//keywords
queries:
	for _, q := range query {

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// find the first word then quit
				continue queries
			}
		}
	}

	//goto
	var i int
increment:
	fmt.Println(i)
	i++
	if i < 5 {
		goto increment
	}
}
