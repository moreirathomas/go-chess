package color

import "fmt"

var (
	Black = colorizeLog("\033[1;30m%s\033[0m")
	White = colorizeLog("\033[1;37m%s\033[0m")
)

func colorizeLog(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}
