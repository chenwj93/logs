package logs

import "strings"

type DEBUGLEVEL int

const (
	DEBUG DEBUGLEVEL = 0
	INFO  DEBUGLEVEL = 1
	WARN  DEBUGLEVEL = 2
	ERROR DEBUGLEVEL = 4
	FATAL DEBUGLEVEL = 8
)

func GetLevel(level string) DEBUGLEVEL {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG" :
		return DEBUG
	case "INFO" :
		return INFO
	case "WARN" :
		return WARN
	case "ERROR" :
		return ERROR
	case "FATAL" :
		return FATAL
	default:
		return WARN
	}
}
