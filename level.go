package firelog

import (
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	DEBUG = uint8(1 << iota)
	TRACE
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

var LevelColor = map[uint8]uint8{
	INFO:  none,
	DEBUG: blue,
	TRACE: green,
	WARN:  yellow,
	ERROR: red,
}

var LevelMap = map[uint8]string{
	DEBUG: "D",
	TRACE: "T",
	INFO:  "I",
	WARN:  "W",
	ERROR: "E",
	FATAL: "F",
	PANIC: "P",
}

type Record struct {
	Time    string  `json:"time,omitempty"`
	Level   uint8   `json:"level,omitempty"`
	Module  string  `json:"module,omitempty"`
	FuncPtr uintptr `json:"func_ptr,omitempty"`
	File    string  `json:"file,omitempty"`
	Line    int     `json:"line,omitempty"`
	Msg     string  `json:"msg,omitempty"`
}

func (rcd *Record) Format(str string) string {
	str = strings.Replace(str, "TIME", time.Now().Format("2021-09-21 22:00:00"), -1)
	str = strings.Replace(str, "LEVEL", LevelMap[rcd.Level], -1)
	str = strings.Replace(str, "MODULE", rcd.Module, -1)
	str = strings.Replace(str, "NCNAME", FuncName(rcd.FuncPtr), -1)
	str = strings.Replace(str, "PATH", path.Dir(rcd.File), -1)
	str = strings.Replace(str, "FILE", path.Base(rcd.File), -1)
	str = strings.Replace(str, "LINE", strconv.Itoa(rcd.Line), -1)
	str = strings.Replace(str, "MSG", rcd.Msg, -1)

	return str
}
