package firelog

import (
	"encoding/json"
)

const (
	// DEFAULTFORMAT default log message format.
	DEFAULTFORMAT = "TIME [LEVEL] FILE:LINE MESSAGE"
	// ModeColor color mode.
	ModeColor = 1 << iota
)

type TextFormat struct {
	format string
	mode   int
}

type JSONFormat struct{}

func (f *JSONFormat) Format(rcd *Record) string {
	b, _ := json.Marshal(rcd)
	return string(b)
}

type Formatter interface {
	Format(rcd *Record) string
}

func NewTextFormat(f string, m int) Formatter {
	return &TextFormat{format: f, mode: m}
}

func (f *TextFormat) Format(rcd *Record) string {
	msg := rcd.Format(f.format)
	if f.mode&ModeColor != 0 {
		return color(LevelColor[rcd.Level], msg)
	}
	return msg
}
