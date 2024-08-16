package config

import (
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

func NewLogger(viper *viper.Viper) *slog.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(viper.GetInt64("LOG_LEVEL")),
	}))

	return log
}
