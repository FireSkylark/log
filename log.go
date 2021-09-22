package firelog

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
)

const (
	ModeSync = 1 << iota
	ModeAsync
)

type Log struct {
	mux        *sync.Mutex
	pool       sync.Pool
	name       string
	mode       int
	level      uint8
	depth      int
	delim      string
	timeFormat string
	Formatter
	out []io.Writer
}

func NewLog(name string, out ...io.Writer) *Log {
	return &Log{
		mux:        new(sync.Mutex),
		pool:       sync.Pool{New: func() interface{} { return new(Record) }},
		name:       name,
		mode:       ModeSync,
		level:      DEBUG,
		depth:      2,
		delim:      "\n",
		timeFormat: time.RFC3339,
		Formatter:  NewTextFormat(DEFAULTFORMAT, 0),
		out:        out,
	}
}

func (l *Log) SetOutput(out io.Writer) {
	l.out = append(l.out, out)
}

func (l *Log) SetTimeFormat(timeFormat string) {
	l.timeFormat = timeFormat
}

func (l *Log) SetFormatter(format Formatter) {
	l.Formatter = format
}

func (l *Log) SetLevel(level uint8) {
	l.level = level
}

func (l *Log) SetMode(mode int) {
	l.mode = mode
}

func (l *Log) SetCallDepth(depth int) {
	l.depth = depth
}

func (l *Log) Output(lv uint8, format string, v ...interface{}) {
	l.mux.Lock()
	defer l.mux.Unlock()
	if lv < l.level {
		return
	}
	ctn := l.pool.Get().(*Record)
	ctn.Time = time.Now().Format(l.timeFormat)
	ctn.Level = lv
	ctn.Module = l.name
	ctn.Msg = fmt.Sprintf(format, v...)
	ctn.FuncPtr, ctn.File, ctn.Line, _ = runtime.Caller(l.depth)

	msg := l.Format(ctn) + l.delim

	for _, out := range l.out {
		_, err := fmt.Fprint(out, msg)
		if err != nil {
			return
		}
	}
	l.pool.Put(ctn)
}

func (l *Log) Trace(format string, v ...interface{}) {
	l.Output(TRACE, format, v...)
}

func (l *Log) Debug(format string, v ...interface{}) {
	l.Output(DEBUG, format, v...)
}

func (l *Log) Info(format string, v ...interface{}) {
	l.Output(INFO, format, v...)
}

func (l *Log) Warn(format string, v ...interface{}) {
	l.Output(WARN, format, v...)
}

func (l *Log) Error(format string, v ...interface{}) {
	l.Output(ERROR, format, v...)
}
