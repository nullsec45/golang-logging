package golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	
	logger.WithField("username","fajar").Info("Hello World")

	logger.WithField("username","fajar").WithField("name","nullsec45").Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username":"nullsec45",
		"name":"Rama Fajar"
	}).Info("Hello World")
}