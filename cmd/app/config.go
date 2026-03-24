package main

import (
	"aurevoir/internal/theme"
)

type Cfg struct {
	Theme theme.Cfg `toml:"theme"`
}
