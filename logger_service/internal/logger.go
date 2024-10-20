package internal

import (
	"log"
	"sync"
)

type Logger struct {
	mu sync.Mutex
}

// NewLogger инициализирует логгер
func NewLogger() *Logger {
	return &Logger{}
}

// LogEvent логирует событие и сохраняет его в MongoDB и RabbitMQ
func (l *Logger) LogEvent(level, message, service string, data map[string]string) {
	log.Printf("[%s] %s: %s - %v", level, service, message, data)
}
