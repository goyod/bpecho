package logger

import (
	"github.com/goyod/pbecho/config"
)

type Logger struct{}

func New(*config.Logger) *Logger {
	return &Logger{}
}
