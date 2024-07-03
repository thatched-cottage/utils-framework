package main

import (
	"github.com/thatched-cottage/utils-framework/pkg/log"
	"testing"
)

func TestLog(t *testing.T) {
	log.Init(log.WithRotatingFileWriter("./log", "test"))
	log.Debug("test log")
}
