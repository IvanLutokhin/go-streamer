package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/IvanLutokhin/go-streamer/internal/app/streamer"
	"github.com/IvanLutokhin/go-streamer/internal/app/streamer/config"
	"github.com/IvanLutokhin/go-streamer/internal/pkg/version"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
)

var (
	versionPtr *bool
	configPtr  *string
)

func init() {
	versionPtr = flag.Bool("version", false, "Display application version")
	configPtr = flag.String("config", "", "Path to a config file. Example is ./configs/streamer.default.json")
}

func main() {
	flag.Parse()

	if *versionPtr {
		fmt.Println(version.String())

		os.Exit(0)
	}

	configPath := os.Getenv("STREAMER_CONFIG")
	if *configPtr != "" {
		configPath = *configPtr
	}

	var configProvider fx.Option
	if configPath != "" {
		c, err := config.Load(configPath)
		if err != nil {
			fmt.Println(err)

			os.Exit(1)
		}

		configProvider = fx.Provide(func() *config.Config { return c })
	} else {
		configProvider = fx.Provide(func() *config.Config { return config.NewDefault() })
	}

	app := streamer.New(configProvider)

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
