package streamer

import (
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
	}

	return append(baseOpts, opts...)
}
