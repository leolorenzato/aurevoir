package main

import (
	"fmt"
	"os"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
)

func getLogFile() (*os.File, error) {
	logDir := filepath.Join(os.Getenv("HOME"), ".config", appName, "log")
	os.MkdirAll(logDir, os.ModePerm)
	logFile := filepath.Join(logDir, "debug.log")
	err := os.Truncate(logFile, 0)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("failed to truncate log file:", err)
		os.Exit(1)
	}

	return tea.LogToFile(logFile, "")
}
