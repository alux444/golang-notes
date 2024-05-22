package main

import (
	"fmt"
	"strconv"
	"time"
)

type digit [5]string

func printTime(time string) {
	zero := digit{
		" ███ ",
		"█   █",
		"█   █",
		"█   █",
		" ███ ",
	}

	one := digit{
		"  █  ",
		" ██  ",
		"  █  ",
		"  █  ",
		" ███ ",
	}

	two := digit{
		" ███ ",
		"█   █",
		"   █ ",
		"  █  ",
		"█████",
	}

	three := digit{
		" ███ ",
		"█   █",
		"  ██ ",
		"█   █",
		" ███ ",
	}

	four := digit{
		"█   █",
		"█   █",
		"█████",
		"    █",
		"    █",
	}

	five := digit{
		"█████",
		"█    ",
		"████ ",
		"    █",
		"████ ",
	}

	six := digit{
		" ███ ",
		"█    ",
		"████ ",
		"█   █",
		" ███ ",
	}

	seven := digit{
		"█████",
		"   █ ",
		"  █  ",
		" █   ",
		"█    ",
	}

	eight := digit{
		" ███ ",
		"█   █",
		" ███ ",
		"█   █",
		" ███ ",
	}

	nine := digit{
		" ███ ",
		"█   █",
		" ████",
		"    █",
		" ███ ",
	}

	colon := digit{
		"     ",
		"  ▒  ",
		"     ",
		"  ▒  ",
		"     ",
	}

	blank := digit{
		"     ",
		"     ",
		"     ",
		"     ",
		"     ",
	}

	digits := [...]digit{zero, one, two, three, four, five, six, seven, eight, nine, colon, blank}

	for i := range digits[0] {
		for j := range len(time) {
			currentChar := time[j]
			currentIndex := currentChar - '0'

			if currentChar == ':' {
				fmt.Print(digits[10][i])
			} else if currentChar == ' ' {
				fmt.Print(digits[11][i])
			} else {
				fmt.Print(digits[currentIndex][i] + " ")
			}
		}
		fmt.Println()
	}
}

func printCurrentTime(showColons bool) {
	time := time.Now()
	hour, minute, second := time.Hour(), time.Minute(), time.Second()

	hourString := strconv.Itoa(hour)
	if hour < 10 {
		hourString = "0" + hourString
	}

	minuteString := strconv.Itoa(minute)
	if minute < 10 {
		minuteString = "0" + minuteString
	}

	secondString := strconv.Itoa(second)
	if second < 10 {
		secondString = "0" + secondString
	}

	var separator string
	if showColons {
		separator = ":"
	} else {
		separator = " "
	}
	timeString := hourString + separator + minuteString + separator + secondString

	printTime(timeString)
}

func main() {
	showColons := true
	for {
		fmt.Print("\033[H\033[2J") // Clear the terminal
		printCurrentTime(showColons)
		showColons = !showColons
		time.Sleep(time.Second)
	}
}
