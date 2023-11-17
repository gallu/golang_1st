package main

/* go 1.21以降じゃないと動かないので注意 */

import (
	"log/slog"
	"os"
)

func main() {
	{
		slog.Info("info")
		slog.Info("hello", "count", 3)
		slog.Info("info",
			slog.Int("num", 123),
			slog.Bool("bool", true),
			slog.String("str", "abc"),
		)
	}

	{
		logger := slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug},
			))
		slog.SetDefault(logger)

		slog.Debug("debug")
		slog.Debug("debug",
			slog.Int("num", 123),
			slog.Bool("bool", true),
			slog.String("str", "abc"),
		)
	}

	{
		logger := slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug},
			))
		slog.SetDefault(logger)

		slog.Debug("debug")
		slog.Debug("debug",
			slog.Int("num", 123),
			slog.Bool("bool", true),
			slog.String("str", "abc"),
		)
	}
}
