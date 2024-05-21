package main

import (
	"fmt"
	"time"
)

func getSeason(args string) {
	switch m := args; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month.\n", m)
	}
}

func getTime() {
	switch h := time.Now().Hour(); {
		case h >= 18: // 18 to 23
			fmt.Println("good evening")
		case h >= 12: // 12 to 18
			fmt.Println("good afternoon")
		case h >= 6: // 6 to 12
			fmt.Println("good morning")
		default: // 0 to 5
			fmt.Println("good night")
		}
}

func main() {
	getSeason("Apr")
	getTime()
}