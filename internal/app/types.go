package app

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

type UserSettings struct {
	studyTime *Timer
	restTime  *Timer
}
