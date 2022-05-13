package core

import (
	"github.com/sirupsen/logrus"
)

// Config controls the customizable behavior of any/all actions.
type Config struct {
	Logger         *logrus.Entry
	RequestsFolder string
}
