package config

import (
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		info:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
		warn:  log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime),
		error: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime),
	}
}

func (l *Logger) Info(msg string) {
	l.info.Println(msg)
}

func (l *Logger) Warn(msg string) {
	l.warn.Println(msg)
}

func (l *Logger) Error(msg string) {
	l.error.Println(msg)
}
