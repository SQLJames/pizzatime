package balldontlie

import (
	"github.com/sirupsen/logrus"
)

// Config controls the customizable behavior of any/all actions.
type Config struct {
	Logger *logrus.Entry
}

func New(logger *logrus.Entry) *Config {
	return &Config{Logger: logger}
}
