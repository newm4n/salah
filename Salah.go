package salah

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

// StackInfo represent a stack information record in a stacktrace
type StackInfo struct {
	FullPath string
	File     string
	Function string
	Line     int
}

// String return the string representation of StackInfo
func (si *StackInfo) String() string {
	var fileName string
	if strings.LastIndex(si.File, "/") >= 0 {
		fileName = si.File[strings.LastIndex(si.File, "/")+1:]
	} else {
		fileName = si.File
	}
	return fmt.Sprintf("%s:%d - %s()", fileName, si.Line, si.Function)
}

var (
	logger logrus.StdLogger
)

func init() {
	logger = logrus.StandardLogger()
}

// PrintStackTrace print the current stacktrace of the caller
func PrintStackTrace() {
	traces := GetStackTrace(3)
	fmt.Println(StackTraceInfo(traces))
}

// StackTraceString return the current caller stacktrace
func StackTraceString() string {
	traces := GetStackTrace(3)
	return StackTraceInfo(traces)
}

// StackTraceInfo translate stack trace info array into string.
func StackTraceInfo(infos []*StackInfo) string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("Stacktrace :\n"))
	for _, t := range infos {
		buff.WriteString(fmt.Sprintf("\t%s\n", t.String()))
	}
	return buff.String()
}

// GetStackTrace construct slice that represent stack trace information
func GetStackTrace(skip int) []*StackInfo {
	ptrs := make([]uintptr, 20)
	size := runtime.Callers(skip, ptrs)
	traces := make([]*StackInfo, 0)
	if size > 0 {
		frames := runtime.CallersFrames(ptrs[:size])
		for {
			frame, more := frames.Next()
			var fileName string
			if strings.LastIndex(frame.File, "/") >= 0 {
				fileName = frame.File[strings.LastIndex(frame.File, "/")+1:]
			} else {
				fileName = frame.File
			}
			st := &StackInfo{
				FullPath: frame.File,
				File:     fileName,
				Function: frame.Function,
				Line:     frame.Line,
			}
			traces = append(traces, st)
			if !more {
				break
			}
		}
	}
	return traces
}
