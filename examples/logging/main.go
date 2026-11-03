package main

import (
	"log/slog"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info(
		"任务完成",
		"task_id", 123,
		"duration", 42*time.Millisecond,
	)

	logger.Warn(
		"任务重试",
		"task_id", 124,
		"attempt", 2,
	)
}
