package firelog

import (
	"runtime"
	"strings"
)

func FuncName(pc uintptr) string {
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "---"
	}

	name := fn.Name()

	if period := strings.LastIndex(name, "."); period >= 0 {
		name = name[period+1:]
	}
	return name
}
