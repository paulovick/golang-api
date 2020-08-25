package customlogger

import (
	"fmt"
	"log"
)

type CustomLogger struct {
	Info func(message ...interface{})
	Fatal func(message ...interface{})
}

func render(level string, section string, message interface{}) {
	log.Println(fmt.Sprintf("[%s] [%s] %s", level, section, message))
}

func GetLogger(section string) *CustomLogger {
	return &CustomLogger{
		func (message ...interface{}) { render("INFO", section, message) },
		func (message ...interface{}) { render("FATAL", section, message) },
	}
}
