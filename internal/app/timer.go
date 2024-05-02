package app

import (
	"strconv"
)

func InitTimer(seconds, minutes, hours int) Timer {
	var timer Timer

	SetTimer(&timer, seconds, minutes, hours)

	return timer
}

func SetTimer(timer *Timer, seconds, minutes, hours int) {
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
}

func DecTimer(timer *Timer) bool {
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

func GetTimerOutput(timer *Timer) string {
	var output string

	if timer.hours == 0 {
		output = ""
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
