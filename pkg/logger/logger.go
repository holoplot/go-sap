package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

func Setup() {
	handler := tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: "15:04:05.000000",
		NoColor:    os.Getenv("NO_COLOR") != "",
		AddSource:  true,
	})

	slog.SetDefault(slog.New(handler))
}
