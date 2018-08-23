package main

import (
	"log"
	"log/syslog"
)

type Logger interface {
	Info(message string)
	Panic(message string)
}

type DefaultLogger struct {}

type SyslogLogger struct {
	Writer *syslog.Writer
}

func CreateSyslogLogger() (*SyslogLogger, error) {
	writer, err := syslog.New(syslog.LOG_LOCAL3, "bot")

	if err != nil {
		return nil, err
	}

	return &SyslogLogger{Writer: writer}, nil
}

func (l DefaultLogger) Info(message string) {
	log.Println(message)
}

func (l DefaultLogger) Panic(message string) {
	log.Panicln(message)
}

func (l SyslogLogger) Info(message string) {
	l.Writer.Info(message)
}

func (l SyslogLogger) Panic(message string) {
	l.Writer.Crit(message)
	panic(message)
}


