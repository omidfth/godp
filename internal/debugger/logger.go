package debugger

import (
	"log"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var debugMode bool
var debugLevel int

func SetDebugMode(state bool, level int) {
	debugMode = state
	debugLevel = level
}

func getColor(level int) string {
	switch level {
	case TRACE:
		return blue
	case DEBUG:
		return cyan
	case INFO:
		return green
	case WARN:
		return yellow
	case ERROR:
		return red
	case FATAL:
		return magenta
	}
	return white
}

func Debug(tag string, level int, v ...interface{}) {
	if debugMode && debugLevel >= level {
		log.Println(getColor(level), tag, reset, v)
	}
}
