package main

import (
	"log"
)

func main() {
	// -> create a channel
	ch := make(chan string)

	// -> create a logger
	logger := NewLoggerLocal(ch)

	// -> run the logger
	go logger.Run()

	// -> log messages
	logger.Log("first message")
	logger.Log("second message")
	logger.Log("third message")

	// -> close the channel
	close(ch)

	log.Println("done")
}


// _____________________________________________________________________________
// Logger: interface that defines the methods that a logger must implement
type Logger interface {
	Run()
	Log(message string)
}

// -> impl: loggerLocal
func NewLoggerLocal(ch chan string) Logger {
	return &loggerLocal{ch: ch}
}

type loggerLocal struct{
	ch chan string
}
func (l *loggerLocal) Run() {
	for {
		select {
		case message, ok := <-l.ch:
			if !ok {
				return
			}
			log.Println("[logger] - ", message)
		}
	}
}
func (l *loggerLocal) Log(message string) {
	l.ch <- message
}
