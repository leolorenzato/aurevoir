package main

import (
	"aurevoir/internal/app"
	"aurevoir/internal/items"
	"aurevoir/internal/theme"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/BurntSushi/toml"
)

const appName string = "aurevoir"

func main() {
	version := flag.Bool("version", false, "print version information")
	cfgPath := flag.String("c", "", "path to configuration file")
	dryRun := flag.Bool("dry-run", false, "dry run mode")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	if *version {
		printVersion()
		os.Exit(0)
	}

	if *debug {
		f, err := getLogFile()
		if err != nil {
			fmt.Printf("failed to setup the logger: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
	} else {
		log.SetOutput(io.Discard)
	}

	cfg := Cfg{
		Items: items.DefaultCfg(),
		Theme: theme.DefaultCfg(),
	}

	var rawCfg RawCfg
	if *cfgPath != "" {
		if _, err := toml.DecodeFile(*cfgPath, &rawCfg); err != nil {
			log.Printf("failed to load configuration file %s", *cfgPath)
			os.Exit(1)
		}

		cfg.Items.MergeRaw(rawCfg.Items)
		if err := cfg.Items.Validate(); err != nil {
			log.Printf("failed to validate the configuration: %v", err)
			os.Exit(1)
		}

		cfg.Theme.MergeRaw(rawCfg.Theme)
	}

	model := app.NewModel(
		appName,
		getEnabledItems(cfg),
		getStyles(cfg),
		*dryRun,
	)
	p := tea.NewProgram(
		model,
	)
	if _, err := p.Run(); err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}
