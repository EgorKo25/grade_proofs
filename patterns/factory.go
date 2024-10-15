package patterns

import (
	"fmt"
	"log/slog"
)

const (
	loggerSlog = iota
	loggerDefault
)

func NewMyLogger() *myLogger { return &myLogger{} }

// myLogger пример структуры, которая имплементирует интерфейс для фабрики
type myLogger struct{}

func (m *myLogger) Info(string, ...any)  {}
func (m *myLogger) Warn(string, ...any)  {}
func (m *myLogger) Error(string, ...any) {}

// ILogger фабричный интерфейс, который имплементируют все объекты фабрики
type ILogger interface {
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
}

// NewLogging фабрика, которая в зависимости от входного значения, выдает объекты разных типов
func NewLogging(loggerType int) (ILogger, error) {
	switch loggerType {
	case loggerSlog:
		return slog.Default(), nil
	case loggerDefault:
		return NewMyLogger(), nil
	default:
		return nil, fmt.Errorf("unknown logger type")
	}
}
