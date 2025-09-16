package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.WithField("name","fajar").Info("Hello Logging")

	entry := logrus.NewEntry(logger)
	entry.WithField("username","nullsec45")
	entry.Info("Hello Entry")
}