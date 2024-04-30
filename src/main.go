package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mattn/go-tty"
)

type Timer struct {
	seconds int
	minutes int
	hours   int
}

type ProgramState struct {
	paused bool
	status string
	timer  *Timer
}

var state = ProgramState{
	status: "study",
	paused: false,
	timer:  nil,
}

func getUserInput(state *ProgramState) {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		char, err := tty.ReadRune()
		if err != nil {
			panic(err)
		}

		switch char {
		case 'p':
			state.paused = true
		case 'r':
			state.paused = false
		case 'q':
			handleExit()
		}

		// Sleep on input to reduce cpu usage
		time.Sleep(5 * time.Millisecond)
	}
}

func printMenu(state *ProgramState) {
	var pause string
	if state.paused {
		pause += "r: Resume"
	} else {
		pause += "p: Pause"
	}
	fmt.Println(pause, "q: Quit")
	fmt.Println(strings.ToUpper(state.status))
}

func handleDone(state *ProgramState) {
	if state.status == "study" {
		setTimer(state.timer, 0, 5, 0)
		state.status = "rest"
	} else {
		setTimer(state.timer, 0, 5, 0)
		state.status = "study"

	}
}

func initPomodoro(state *ProgramState) {
	initCleanup()
	var timer Timer
	setTimer(&timer, 0, 1, 0)

	state.timer = &timer
	go getUserInput(state)
}

func timerLoop(state *ProgramState) {
	for {
		newLoopPrint()
		printMenu(state)
		if state.paused {
			output := getTimerOutput(state.timer)

			fmt.Println(output)
			// Sleep on pause to reduce cpu usage
			time.Sleep(1 * time.Second)
			continue
		}

		done := decTimer(state.timer)

		output := getTimerOutput(state.timer)

		fmt.Println(output)

		if done {
			handleDone(state)
			continue
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	initPomodoro(&state)

	timerLoop(&state)
}
