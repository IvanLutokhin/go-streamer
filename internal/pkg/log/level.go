package log

import (
	"errors"
	"strings"
)

//go:generate stringer -type=Level

type Level uint8

const (
	UNKNOWN Level = iota
	EMERGENCY
	ALERT
	CRITICAL
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG
)

func LevelFromString(value string) (Level, error) {
	for i := 1; i < len(_Level_index)-1; i++ {
		if strings.EqualFold(value, _Level_name[_Level_index[i]:_Level_index[i+1]]) {
			return Level(i), nil
		}
	}

	return UNKNOWN, errors.New("invalid log level")
}
