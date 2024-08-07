package structural

import "fmt"

type oldLogger interface {
	log(message string)
}

type newLogger interface {
	info(message string)
	error(message string)
	debug(message string)
}

type myNewLogger struct{}

func (n *myNewLogger) info(message string) {
	fmt.Println("info: ", message)
}

func (n *myNewLogger) error(message string) {
	fmt.Println("error: ", message)
}

func (n *myNewLogger) debug(message string) {
	fmt.Println("debug: ", message)
}

type loggerAdapter struct {
	logger newLogger
}

func NewLoggerAdapter(logger newLogger) oldLogger {
	return &loggerAdapter{logger: logger}
}

func (adapter *loggerAdapter) log(message string) {
	adapter.logger.info(message)
}

func runAdapter() {
	newLogger := &myNewLogger{}
	logger := NewLoggerAdapter(newLogger)
	logger.log("abc")

}
