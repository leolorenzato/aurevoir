package main

import (
	"aurevoir/internal/items"
	"aurevoir/internal/theme"
)

type RawCfg struct {
	Items items.RawCfg `toml:"items"`
	Theme theme.RawCfg `toml:"theme"`
}

type Cfg struct {
	Items items.Cfg
	Theme theme.Cfg
}
