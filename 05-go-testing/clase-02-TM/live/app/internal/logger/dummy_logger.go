package logger

// constructor
func NewLoggerDummy() *LoggerDummy {
	return &LoggerDummy{}
}

// LoggerDummy implements the Logger interface.
type LoggerDummy struct {}

func (l *LoggerDummy) Log(msg string) {
	// do nothing
}