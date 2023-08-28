package main

import "golang.org/x/exp/slog"

func main() {
	slog.Info("Started", "count", 3)
}
