package agent

import (
	"fmt"
	// "os"
	"runtime"
	"strings"
	"time"
)

const (
	Gray = uint8(iota + 90)
	Red
	Green
	Yellow
	Blue
	Magenta
	EndColor = "\033[0m"
	INFO     = "INFO"
	TRAC     = "TRAC"
	ERROR    = "ERROR"
	WARN     = "WARN"
	SUCC     = "SUCC"
)

func ColorLog(format string, a ...interface{}) {
	fmt.Print(ColorLogS(format, a...))
}

func ColorLogS(format string, a ...interface{}) string {
	log := fmt.Sprintf(format, a...)
	var clog string
	if runtime.GOOS != "windows" {
		// Level.
		i := strings.Index(log, "]")
		if log[0] == '[' && i > -1 {
			clog += "[" + getColorLevel(log[1:i]) + "]"
		}

		log = log[i+1:]

		// Error.
		log = strings.Replace(log, "[ ", fmt.Sprintf("[\033[%dm", Red), -1)
		log = strings.Replace(log, " ]", EndColor+"]", -1)

		// Path.
		log = strings.Replace(log, "( ", fmt.Sprintf("(\033[%dm", Yellow), -1)
		log = strings.Replace(log, " )", EndColor+")", -1)

		// Highlights.
		log = strings.Replace(log, "# ", fmt.Sprintf("\033[%dm", Gray), -1)
		log = strings.Replace(log, " #", EndColor, -1)

		log = clog + log

	} else {
		// Level.
		i := strings.Index(log, "]")
		if log[0] == '[' && i > -1 {
			clog += "[" + log[1:i] + "]"
		}

		log = log[i+1:]

		// Error.
		log = strings.Replace(log, "[ ", "[", -1)
		log = strings.Replace(log, " ]", "]", -1)

		// Path.
		log = strings.Replace(log, "( ", "(", -1)
		log = strings.Replace(log, " )", ")", -1)

		// Highlights.
		log = strings.Replace(log, "# ", "", -1)
		log = strings.Replace(log, " #", "", -1)

		log = clog + log
	}

	return time.Now().Format("2006/01/02 15:04:05 ") + log
}

func getColorLevel(level string) string {
	level = strings.ToUpper(level)
	switch level {
	case INFO:
		return fmt.Sprintf("\033[%dm%s\033[0m", Blue, level)
	case TRAC:
		return fmt.Sprintf("\033[%dm%s\033[0m", Blue, level)
	case ERROR:
		return fmt.Sprintf("\033[%dm%s\033[0m", Red, level)
	case WARN:
		return fmt.Sprintf("\033[%dm%s\033[0m", Magenta, level)
	case SUCC:
		return fmt.Sprintf("\033[%dm%s\033[0m", Green, level)
	default:
		return level
	}
}

// func WriteLogToFile(filename, taskname, logstring string) bool {
// 	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeType)
// 	if err != nil {
// 		return false
// 	}
// 	defer file.Close()
// 	_, err = file.WriteString(ColorLog("[INFO] 执行任务: %v 输出: %s\n", taskname, logstring))
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }
