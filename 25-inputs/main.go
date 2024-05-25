package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func inputs() {
	// Simulate an error
	// os.Stdin.Close()

	// Create a new scanner that scans from the standard-input
	in := bufio.NewScanner(os.Stdin)

	// Stores the total number of lines in the input
	var lines int

	// Scan the input line by line
	for in.Scan() {
		lines++

		fmt.Println("Scanned Text :", in.Text())
		fmt.Println("Scanned Bytes:", in.Bytes())
		in.Text()
		if in.Text() == "end" {
			break
		}
	}

	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}
}

func uppercaser() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		if in.Text() == "end" {
			break
		}
		fmt.Println(strings.ToUpper(in.Text()))
	}
}

func uniqueWords() {
	in := bufio.NewScanner(os.Stdin)
	rx := regexp.MustCompile(`[^A-Za-z]+`)

	total, words := 0, make(map[string]int)
	for in.Scan() {
		total++

		word := rx.ReplaceAllString(in.Text(), "")
		word = strings.ToLower(word)
		words[word]++
	}

	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))
}

func logparser() {
	//go run main.go < log.txt
	dict := make(map[string]int)

	in := bufio.NewScanner(os.Stdin)
	lines := 0

	for in.Scan() {
		lines++
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("Incorrect input at line %d\n", lines)
			return
		}
		integer, error := strconv.Atoi(fields[1])
		if error != nil {
			fmt.Printf("Error parsing integer value at line %d\n", lines)
			return
		}
		if integer < 0 {
			fmt.Printf("Negative integer value at line %d\n", lines)
			return
		}

		dict[fields[0]] = dict[fields[0]] + integer
	}

	for key, value := range dict {
		fmt.Printf("%s had %d visitors.\n", key, value)
	}
}

func main() {
	logparser()
}
