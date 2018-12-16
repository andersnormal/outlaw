package cmd

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

type root struct {
	cancel context.CancelFunc
	logger *log.Entry
	ctx    context.Context
	sys    chan os.Signal
	exit   chan int
}
