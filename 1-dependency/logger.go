package main

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type Logger interface {
	logrus.FieldLogger
}

func NewLogger(lc fx.Lifecycle) Logger {
	logger := logrus.New()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting logger")
			return nil
		},
	})
	return logger
}
