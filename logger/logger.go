//package logger helps in logging .
// in Envmode == live it print all logs in syslog
// otheriwse use default go's log mechanism
package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/blackjack/syslog"
	"gitlab.com/varver/wmd/config"
)

const (
	Dev  = "dev"
	Live = "live"
)

func init() {
	progname := os.Args[0]
	base := path.Base(progname)
	syslog.Openlog(base, syslog.LOG_PID, syslog.LOG_USER)
}

func llogf(p syslog.Priority, format string, a ...interface{}) {
	_, f, l, _ := runtime.Caller(2)
	fd := strings.Split(f, "/src/")
	if len(fd) == 2 {
		filedata := fmt.Sprintf("%s.%d: ", fd[1], l)
		format = filedata + format
	}
	syslog.Syslogf(p, format, a)
}
func llog(p syslog.Priority, msg string) {
	_, f, l, _ := runtime.Caller(2)
	fd := strings.Split(f, "/src/")
	if len(fd) == 2 {
		filedata := fmt.Sprintf("%s.%d: ", fd[1], l)
		msg = filedata + msg
	}
	syslog.Syslog(p, msg)
}
func Emerg(msg string) {
	msg = "Emerg : " + msg
	if config.Setting.EnvMode == Dev {
		log.Println(msg)
		return
	}
	llog(syslog.LOG_EMERG, msg)
}
func Emergf(format string, a ...interface{}) {
	format = "Emerg : " + format
	if config.Setting.EnvMode == Dev {
		log.Printf(format, a)
		return
	}
	llogf(syslog.LOG_EMERG, format, a...)
}
func Alert(msg string) {
	msg = "Alert : " + msg
	if config.Setting.EnvMode == Dev {
		log.Println(msg)
		return
	}
	llog(syslog.LOG_ALERT, msg)
}
func Alertf(format string, a ...interface{}) {
	format = "Alert : " + format
	if config.Setting.EnvMode == Dev {
		log.Printf(format, a)
		return
	}
	llogf(syslog.LOG_ALERT, format, a...)
}
func Crit(msg string) {
	msg = "Crit : " + msg
	if config.Setting.EnvMode == Dev {
		log.Println(msg)
		return
	}
	llog(syslog.LOG_CRIT, msg)
}
func Critf(format string, a ...interface{}) {
	format = "Crit : " + format
	if config.Setting.EnvMode == Dev {
		log.Printf(format, a)
		return
	}
	llogf(syslog.LOG_CRIT, format, a...)
}
func Err(msg string) {
	msg = "Err : " + msg
	if config.Setting.EnvMode == Dev {
		log.Println(msg)
		return
	}
	llog(syslog.LOG_ERR, msg)
}
func Errf(format string, a ...interface{}) {
	format = "Err : " + format
	if config.Setting.EnvMode == Dev {
		log.Printf(format, a)
		return
	}
	llogf(syslog.LOG_ERR, format, a...)
}
func Warning(msg string) {
	msg = "Warning : " + msg
	if config.Setting.EnvMode == Dev {
		log.Println(msg)
		return
	}
	llog(syslog.LOG_WARNING, msg)
}
func Warningf(format string, a ...interface{}) {
	format = "Warning : " + format
	if config.Setting.EnvMode == Dev {
		log.Printf(format, a)
		return
	}
	llogf(syslog.LOG_WARNING, format, a...)
}
func Notice(msg string) {
	msg = "Notice : " + msg
	if config.Setting.EnvMode == Dev {
		log.Println(msg)
		return
	}
	llog(syslog.LOG_NOTICE, msg)
}
func Noticef(format string, a ...interface{}) {
	format = "Notice : " + format
	if config.Setting.EnvMode == Dev {
		log.Printf(format, a)
		return
	}
	llogf(syslog.LOG_NOTICE, format, a...)
}
func Info(msg string) {
	msg = "Info : " + msg
	if config.Setting.EnvMode == Dev {
		log.Println(msg)
		return
	}
	llog(syslog.LOG_INFO, msg)
}
func Infof(format string, a ...interface{}) {
	format = "Info : " + format
	if config.Setting.EnvMode == Dev {
		log.Printf(format, a)
		return
	}
	llogf(syslog.LOG_INFO, format, a...)
}
func Debug(msg string) {
	msg = "Debug : " + msg
	if config.Setting.EnvMode == Dev {
		log.Println(msg)
		return
	}
	llog(syslog.LOG_DEBUG, msg)
}
func Debugf(format string, a ...interface{}) {
	format = "Debug : " + format
	if config.Setting.EnvMode == Dev {
		log.Printf(format, a)
		return
	}
	llogf(syslog.LOG_DEBUG, format, a...)
}
