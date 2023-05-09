package logger

// Logger is the interface that wraps the basic logging methods.
type Logger interface {
	// Log logs a message at the Info level. (async task)
	Log(msg string)
}