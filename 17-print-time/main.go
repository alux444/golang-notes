package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func printTime(time string) {
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

func printCurrentTime(showColons bool, leftShift int) {
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

	if leftShift > 0 && leftShift < 8 {
		timeString = timeString[leftShift:] + strings.Repeat(" ", leftShift)
	}

	if leftShift >= 8 {
		repeatTimes := 16 - leftShift
		endIndex := leftShift - 8
		timeString = strings.Repeat(" ", repeatTimes) + timeString[:endIndex]
	}

	printTime(timeString)
}

func main() {
	showColons := true
	leftShift := 0
	for {
		fmt.Print("\033[H\033[2J") // Clear the terminal

		printCurrentTime(showColons, leftShift)
		showColons = !showColons
		leftShift++
		if leftShift == 16 {
			leftShift = 0
		}
		time.Sleep(time.Second)
	}
}
