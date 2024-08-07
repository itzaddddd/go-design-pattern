package behavior

import "fmt"

type Logger interface {
	setNextLogger(next Logger)
	logMessage(level int, message string)
	write(message string)
}

type DefaultLogger struct {
	level      int
	nextLogger Logger
}

func NewDefaultLogger(level int) *DefaultLogger {
	return &DefaultLogger{
		level: level,
	}
}

func (d *DefaultLogger) setNextLogger(next Logger) {
	d.nextLogger = next
}

func (d *DefaultLogger) logMessage(level int, message string) {
	if d.level <= level {
		d.write(message)
	}
	if d.nextLogger != nil {
		d.nextLogger.logMessage(level, message)
	}
}

func (d *DefaultLogger) write(message string) {}

type ConsoleLogger struct {
	DefaultLogger
}

func NewConsoleLogger(level int) *ConsoleLogger {
	return &ConsoleLogger{DefaultLogger: *NewDefaultLogger(level)}
}

func (c *ConsoleLogger) write(message string) {
	fmt.Println("Standard Console::Logger: " + message)
}

type ErrorLogger struct {
	DefaultLogger
}

func NewErrorLogger(level int) *ErrorLogger {
	return &ErrorLogger{DefaultLogger: *NewDefaultLogger(level)}
}

func (e *ErrorLogger) write(message string) {
	fmt.Println("Error Console::Logger: " + message)
}

type FileLogger struct {
	DefaultLogger
}

func NewFileLogger(level int) *FileLogger {
	return &FileLogger{DefaultLogger: *NewDefaultLogger(level)}
}

func (f *FileLogger) write(message string) {
	fmt.Println("File::Logger: " + message)
}

func RunChain() {
	loggerChain := NewConsoleLogger(1)
	errorLogger := NewErrorLogger(2)
	fileLogger := NewFileLogger(3)

	loggerChain.setNextLogger(errorLogger)
	errorLogger.setNextLogger(fileLogger)

	loggerChain.logMessage(1, "Information")
	loggerChain.logMessage(2, "Error")
	loggerChain.logMessage(3, "Debug message")
}
