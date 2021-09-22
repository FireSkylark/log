package firelog

import (
	"fmt"
)

const (
	red = uint8(iota + 91)
	green
	yellow
	blue
	none = uint8(0)
)

func color(col uint8, s interface{}) string {
	return fmt.Sprintf("\\x1b[%dm%v\\x1b[0m", col, s)
}
