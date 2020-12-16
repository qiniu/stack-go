package log

import "fmt"

// A simple logger implimentation, convenient for tests
type simpleLog struct{}

// NewSimpleLog new a simple logger implimentation, convenient for tests
func NewSimpleLog() Logger {
	return &simpleLog{}
}

func (l *simpleLog) Debugf(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format, args...)
}

func (l *simpleLog) Infof(format string, args ...interface{}) {
	fmt.Printf("[INFO] "+format, args...)
}

func (l *simpleLog) Warnf(format string, args ...interface{}) {
	fmt.Printf("[WARN] "+format, args...)
}

func (l *simpleLog) Errorf(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format, args...)
}
