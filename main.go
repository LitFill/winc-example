// WinUi, LitFill <author at email dot com>
// program for ...
package main

import (
	"github.com/LitFill/fatal"
)

func main() {
	logFile := fatal.CreateLogFile("log.json")
	defer logFile.Close()
	// logger := fatal.CreateLogger(logFile, slog.LevelInfo)

	runGui()
}
