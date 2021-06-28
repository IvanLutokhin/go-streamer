package streamer

import (
	"github.com/IvanLutokhin/go-streamer/internal/app/streamer/logger"
	"go.uber.org/fx"
	"time"
)

func New(opts ...fx.Option) *fx.App {
	return fx.New(buildOpts(opts...)...)
}

func buildOpts(opts ...fx.Option) []fx.Option {
	baseOpts := []fx.Option{
		fx.Options(
			fx.NopLogger,
			fx.StartTimeout(time.Minute),
			fx.StopTimeout(time.Minute),
		),
		fx.Provide(
			logger.New,
		),
	}

	return append(baseOpts, opts...)
}
