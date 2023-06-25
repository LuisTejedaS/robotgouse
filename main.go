package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		robotgo.MoveSmoothRelative(20, 20)
		robotgo.MoveSmoothRelative(-20, -20)
		time.Sleep(time.Minute * 4)
	}
}
