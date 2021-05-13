package main

import (
	"context"
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/app/streamer"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := streamer.New()

	startCtx, startCancel := context.WithTimeout(context.Background(), app.StartTimeout())
	defer startCancel()

	if err := app.Start(startCtx); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	terminate := make(chan os.Signal, 1)

	signal.Notify(terminate, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)

	<-terminate

	stopCtx, stopCancel := context.WithTimeout(context.Background(), app.StopTimeout())
	defer stopCancel()

	if err := app.Stop(stopCtx); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}

	os.Exit(0)
}
