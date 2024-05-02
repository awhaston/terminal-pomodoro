package app

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mattn/go-tty"
)

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
		SetTimer(state.timer, 0, 5, 0)
		state.status = "rest"
	} else {
		SetTimer(state.timer, 0, 5, 0)
		state.status = "study"

	}
}

func PomodoroInit() (*ProgramState, *UserSettings) {
	InitCleanup()
	state := ProgramState{
		status: "study",
		paused: false,
		timer:  nil,
	}

	settings := UserSettings{
		studyTime: nil,
		restTime:  nil,
	}

	// TODO Add user config file and get this values from there
	studyTime := Timer{
		seconds: 0,
		minutes: 1,
		hours:   0,
	}

	restTime := Timer{
		seconds: 30,
		minutes: 0,
		hours:   0,
	}

	settings.studyTime = &studyTime
	settings.restTime = &restTime

	if state.timer == nil {
		var timer Timer
		if state.status == "study" {
			SetTimer(&timer, settings.studyTime.seconds, settings.studyTime.minutes, settings.studyTime.hours)
		} else if state.status == "rest" {
			SetTimer(&timer, settings.restTime.seconds, settings.restTime.minutes, settings.restTime.hours)
		}
		state.timer = &timer
	}

	return &state, &settings
}

func PomodoroLoop(state *ProgramState, settings *UserSettings) {
	go getUserInput(state)
	for {
		newLoopPrint()
		printMenu(state)
		if state.paused {
			output := GetTimerOutput(state.timer)

			fmt.Println(output)
			// Sleep on pause to reduce cpu usage
			time.Sleep(1 * time.Second)
			continue
		}

		done := DecTimer(state.timer)

		output := GetTimerOutput(state.timer)

		fmt.Println(output)

		if done {
			handleDone(state)
			continue
		}

		time.Sleep(1 * time.Second)
	}
}
