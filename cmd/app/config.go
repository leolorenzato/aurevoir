package main

import (
	"aurevoir/internal/items"
	"aurevoir/internal/theme"
)

type Cfg struct {
	Items items.Cfg `toml:"items"`
	Theme theme.Cfg `toml:"theme"`
}
