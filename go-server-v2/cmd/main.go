package main

import (
	"context"
	"os"
	"syscall"

	"go-server-v2/config"
	"go-server-v2/http"
	"go-server-v2/repo"

	"github.com/jesperkha/notifier"
)

func main() {
	config := config.Load()
	notif := notifier.New()

	logger := repo.NewLogger(config.Environment)

	deps := http.Dependencies{
		Logger: logger,
		Notif:  notif,
		Config: config,
	}

	go http.Run(deps)

	notif.NotifyOnSignal(os.Interrupt, syscall.SIGTERM)
	logger.Info(context.Background(), "shutdown complete")
}
