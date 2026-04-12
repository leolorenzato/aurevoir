package main

import (
	"aurevoir/internal/items"
	"aurevoir/internal/theme"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"

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

func getEnabledItems(cfg Cfg) []items.Item {
	var enabledItems []items.Item
	for _, item := range items.Build(cfg.Items) {
		if item.Enable {
			enabledItems = append(enabledItems, item)
		}
	}

	return enabledItems
}

func getStyles(cfg Cfg) theme.Styles {
	return theme.Build(cfg.Theme)
}

func printVersion() {
	version := "unknown"
	commit := "unknown"
	buildTime := "unknown"
	goVersion := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok {
		version = info.Main.Version
		goVersion = info.GoVersion
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				commit = setting.Value
			case "vcs.time":
				buildTime = setting.Value
			}
		}
	}
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Printf("Build time: %s\n", buildTime)
	fmt.Printf("Go version: %s\n", goVersion)
}
