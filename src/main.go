package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

type Timer struct {
	seconds int
	minutes int
	hours   int
}

var clear map[string]func() // create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) // Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") // Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] // runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          // if we defined a clear func for that platform:
		value() // we execute it
	} else { // unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func initTimer(seconds, minutes, hours int) Timer {
	var timer Timer

	if seconds == 0 {
		timer.seconds = 0
	} else {
		timer.seconds = seconds
	}

	if minutes == 0 {
		timer.minutes = 25
	} else {
		timer.minutes = minutes
	}
	if hours == 0 {
		timer.hours = 0
	} else {
		timer.hours = hours
	}

	return timer
}

func decTimer(timer *Timer) bool {
	var done bool

	if timer.seconds-1 < 0 {
		if timer.minutes-1 < 0 {
			if timer.hours-1 < 0 {
				done = true
			} else {
				timer.seconds = 59
				timer.minutes = 59
				timer.hours = timer.hours - 1
			}
		} else {
			timer.seconds = 59
			timer.minutes = timer.minutes - 1
		}
	} else {
		timer.seconds = timer.seconds - 1
	}

	return done
}

func getOutput(timer *Timer) string {
	var output string

	if timer.hours == 0 {
		output = "  "
	} else if timer.hours >= 10 {
		intString := strconv.FormatInt(int64(timer.hours), 10)
		output = intString
		output += " : "
	} else {
		intString := strconv.FormatInt(int64(timer.hours), 10)
		output = "0" + intString
		output += " : "
	}
	if timer.minutes == 0 {
		output += "00"
	} else if timer.minutes >= 10 {
		intString := strconv.FormatInt(int64(timer.minutes), 10)
		output += intString
	} else {
		intString := strconv.FormatInt(int64(timer.minutes), 10)
		output += "0" + intString
	}
	output += " : "
	if timer.seconds == 0 {
		output += "00"
	} else if timer.seconds >= 10 {
		intString := strconv.FormatInt(int64(timer.seconds), 10)
		output += intString
	} else {
		intString := strconv.FormatInt(int64(timer.seconds), 10)
		output += "0" + intString
	}

	return output
}

func main() {
	timer := initTimer(0, 2, 0)

	for {
		fmt.Print("\033[2J\033[H")

		done := decTimer(&timer)

		output := getOutput(&timer)

		fmt.Println(output)

		if done {
			break
		}

		time.Sleep(1 * time.Second)
	}
}
