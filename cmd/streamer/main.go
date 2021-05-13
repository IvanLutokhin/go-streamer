package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/app/streamer"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/version"
	"os"
	"os/signal"
	"syscall"
)

var versionPtr *bool

func init() {
	versionPtr = flag.Bool("version", false, "Display application version")
}

func main() {
	flag.Parse()

	if *versionPtr {
		fmt.Println(version.String())

		os.Exit(0)
	}

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
