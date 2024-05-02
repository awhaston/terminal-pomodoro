package main

import (
	. "pomodoro/internal/app"
)

func main() {
	// Really hate returning the state and settings pointers to put into loop
	// but it seems like the only way to avoid global state
	state, settings := PomodoroInit()

	PomodoroLoop(state, settings)
}
