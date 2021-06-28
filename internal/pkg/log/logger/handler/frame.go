package handler

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

const unknownFile = "unknown"

type Frame struct {
	File string
	Line int
}

func (frame Frame) String() string {
	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
}

func Caller(skip int) Frame {
	pc, file, line, ok := runtime.Caller(skip + 1)
	if pc == 0 || !ok {
		return Frame{File: unknownFile, Line: 0}
	}

	relativeFile, err := relativeFile(file)
	if err != nil {
		return Frame{File: unknownFile, Line: 0}
	}

	return Frame{File: relativeFile, Line: line}
}

func relativeFile(path string) (string, error) {
	src, err := os.Getwd()
	if err != nil {
		return path, err
	}

	i := strings.LastIndex(path, src)
	if i >= 0 {
		return path[i+len(src):], nil
	}

	return path, nil
}
