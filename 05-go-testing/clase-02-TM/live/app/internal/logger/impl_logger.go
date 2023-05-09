package logger

import "log"

// constructor
func NewLogger() Logger {
	return &LoggerLocal{}
}

// implLogger implements the Logger interface.
type LoggerLocal struct {}
// Log logs a message at the Info level. (async task)
func (l *LoggerLocal) Log(msg string) {
	log.Println(msg)
}