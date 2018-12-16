package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/andersnormal/outlaw/config"
)

func (r *root) configureSignals(cfg *config.Config) {
	r.sys = make(chan os.Signal, 1)

	signal.Notify(r.sys, cfg.ReloadSignal, cfg.KillSignal, cfg.TermSignal)
}

func (r *root) exitSignal() {
	r.exit <- 1
}

// watchSignals is watching configured signals
func (r *root) watchSignals(cfg *config.Config) {
	// defer
	defer r.exitSignal()

	// config singals
	r.configureSignals(cfg)

	// loop blocking
	for {
		sig := <-r.sys
		switch sig {
		case syscall.SIGUSR1:
		default:
			r.logger.Info("Gracefully shutdown ...")

			r.cancel()
			return
		}
	}
}
