package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 20

		// drawing buffer length
		// *2 for extra spaces
		// +1 for newlines
		bufLen = (width*2 + 1) * height
	)

	var (
		px, py       int    // ball position
		vx, vy       = 1, 1 // velocities
		prevx, prevy int

		cell rune // current cell (for caching)
	)

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// create a drawing buffer
	buf := make([]rune, 0, bufLen)

	// clear the screen once
	fmt.Print("\033[H\033[2J") // Clear the terminal

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		// remove the previous ball
		board[prevx][prevy] = false

		// put the new ball
		board[px][py] = true

		// rewind the buffer (allow appending from the beginning)
		buf = buf[:0]

		// draw the board into the buffer
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		// print the buffer
		fmt.Print("\033[H\033[2J") // Clear the terminal
		fmt.Print(string(buf))

		// remember the previous ball position
		prevx, prevy = px, py

		// slow down the animation
		time.Sleep(speed)
	}
}
