package pattern

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Logger interface {
	Output(maxdepth int, s string) error
}

const (
	LOG_DEBUG = 1
	LOG_INFO  = 2
	LOG_WARN  = 3
	LOG_ERROR = 4
	LOG_FATAL = 5
)

type NSQLookupd struct {
	logLevel int
	LogLevel string
	Verbose  bool `flag:"verbose"`
	Logger   Logger
}

func NewNSQLookupd(l string) *NSQLookupd {
	n := new(NSQLookupd)
	n.LogLevel = l
	n.logLevel = n.logLevelFromString(l)
	if n.Logger == nil {
		n.Logger = log.New(os.Stderr, "[nsqlookupd]", log.Ldate|log.Ltime|log.Lmicroseconds)
	}
	return n
}

func (n *NSQLookupd) logLevelFromString(level string) int {
	// check log-level is valid and translate to int
	switch strings.ToLower(level) {
	case "debug":
		return LOG_DEBUG
	case "info":
		return LOG_INFO
	case "warn":
		return LOG_WARN
	case "error":
		return LOG_ERROR
	case "fatal":
		return LOG_FATAL
	default:
		return -1
	}
}

func (n *NSQLookupd) logf(level int, f string, args ...interface{}) {
	levelString := "INFO"
	switch level {
	case LOG_DEBUG:
		levelString = "DEBUG"
	case LOG_INFO:
		levelString = "INFO"
	case LOG_WARN:
		levelString = "WARNING"
	case LOG_ERROR:
		levelString = "ERROR"
	case LOG_FATAL:
		levelString = "FATAL"
	}

	if level >= n.logLevel || n.Verbose {
		n.Logger.Output(2, fmt.Sprintf(levelString+": "+f, args...))
	}
}
