package utarrow

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

const (
	// Log Level
	Fatal = iota
	Error
	Warn
	Info
	Debug
)

type Ulog struct {
	log   *log.Logger
	file  *os.File
	level int
}

var (
	defaultLogLevel = Debug
	DefaultLogger   = &Ulog{log: log.New(os.Stdout, "", log.LstdFlags), file: nil, level: defaultLogLevel}
	errLevels       = []string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG"}
)

func New(file string, levelStr string) (*Ulog, error) {
	level := defaultLogLevel
	for lv, str := range errLevels {
		if str == levelStr {
			level = lv
		}
	}
	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return DefaultLogger, err
		}
		logger := log.New(f, "", log.LstdFlags)
		return &Ulog{log: logger, file: f, level: level}, nil
	}
	return &Ulog{log: log.New(os.Stdout, "", log.LstdFlags), file: nil, level: level}, nil
}

func (l *Ulog) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

func (l *Ulog) Error(format string, args ...interface{}) {
	if l.level >= Error {
		l.logCaller(Error, format, args...)
	}
}
func (l *Ulog) Errorln(args ...interface{}) {
	if l.level >= Error {
		l.logCallerln(Error, args...)
	}
}

func (l *Ulog) Debug(format string, args ...interface{}) {
	if l.level >= Debug {
		l.logCaller(Debug, format, args...)
	}
}
func (l *Ulog) Debugln(args ...interface{}) {
	if l.level >= Debug {
		l.logCallerln(Debug, args...)
	}
}

func (l *Ulog) Info(format string, args ...interface{}) {
	if l.level >= Info {
		l.logCaller(Info, format, args...)
	}
}
func (l *Ulog) Infoln(args ...interface{}) {
	if l.level >= Info {
		l.logCallerln(Info, args...)
	}
}

func (l *Ulog) Warn(format string, args ...interface{}) {
	if l.level >= Warn {
		l.logCaller(Warn, format, args...)
	}
}
func (l *Ulog) Warnln(args ...interface{}) {
	if l.level >= Warn {
		l.logCallerln(Warn, args...)
	}
}

func (l *Ulog) Fatal(format string, args ...interface{}) {
	if l.level >= Fatal {
		l.logCaller(Fatal, format, args...)
	}
}
func (l *Ulog) Fatalln(args ...interface{}) {
	if l.level >= Fatal {
		l.logCallerln(Fatal, args...)
	}
}

func (l *Ulog) logCaller(level int, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "callFail"
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short

	l.log.Print(fmt.Sprintf("[%s] %s:%d %s", errLevels[level], file, line, fmt.Sprintf(format, args...)))
}

func (l *Ulog) logCallerln(level int, args ...interface{}) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "callFail"
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short

	l.log.Print(fmt.Sprintf("[%s] [%s]:%d:%s", errLevels[level], file, line, fmt.Sprintln(args...)))
}
