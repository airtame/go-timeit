package timeit

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// EnterResult is the result of the enter function
type enterResult struct{}

// G is the global timer object
var G = New()

// e is a constant
var e = enterResult{}

// Timer is a running structure
type Timer struct {
	start     time.Time
	prev      time.Time
	callDepth int
}

// New returns a timer
func New() Timer {
	now := time.Now()
	t := Timer{
		start: now,
		prev:  now,
	}
	t.log("initialize", 3)
	return t
}

// Trace reports entry and exit times for a function
func (t *Timer) Trace() func() {
	e := t.enter()
	return func() {
		t.exit(e)
	}
}

// Print prints the time since the last measurement
func (t *Timer) Print(format string, msgs ...interface{}) {
	t.log(fmt.Sprintf(format, msgs...), 3)
}

func (t *Timer) enter() enterResult {
	t.callDepth++
	t.log("enter", 4)
	return e
}

func (t *Timer) exit(enterResult) {
	defer func() { t.callDepth-- }()
	t.log("exit", 4)
}

func (t *Timer) log(logMsg string, level int) {
	fileName, lineNum, funcName := callerInfo(level)
	prev := t.prev
	t.prev = time.Now()
	fmt.Printf("%-120s %d %10s %10s %s %-40s\n",
		fmt.Sprintf("%v:%d::%v",
			fileName,
			lineNum,
			funcName),
		time.Now().UnixNano()/1e6,
		fmt.Sprintf("%v ms", toMsecs(time.Since(t.start))),
		fmt.Sprintf("%v ms", toMsecs(t.prev.Sub(prev))),
		strings.Repeat(" ", t.callDepth),
		logMsg,
	)
}

func toMsecs(d time.Duration) int64 {
	return int64(d / time.Millisecond)
}

func callerInfo(level int) (string, int, string) {
	pc, _, _, _ := runtime.Caller(level)
	funcObj := runtime.FuncForPC(pc)

	fileName, lineNum := funcObj.FileLine(pc)
	funcName := funcObj.Name()

	return fileName, lineNum, funcName
}
