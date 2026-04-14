# aurevoir 🌙

![GitHub License](https://img.shields.io/github/license/leolorenzato/aurevoir)
[![Go Report Card](https://goreportcard.com/badge/github.com/leolorenzato/aurevoir)](https://goreportcard.com/report/github.com/leolorenzato/aurevoir)
![GitHub Release](https://img.shields.io/github/v/release/leolorenzato/aurevoir)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/leolorenzato/aurevoir)

**aurevoir** is a lightweight **TUI power menu** written in Go with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

### ✨ Features
- ⚙️ Configurable commands
- 🎨 Theme customization

### 🔧 Build from Source
- check Go version (requires **Go 1.25+**)
    ```bash
    go version
    ```

- build the app
    ```bash
    git clone git@github.com:leolorenzato/aurevoir.git
    cd ./aurevoir
    go build -o ./bin/aurevoir ./cmd/app
    ```

### 🚀 Run
- run the app
    ```bash
    ./bin/aurevoir
    ```

## 🧩 Configuration

### Default
See the [configuration template](https://github.com/leolorenzato/aurevoir/blob/main/config.template.toml)

### Custom
- create a config file (e.g `~/.config/aurevoir/config.toml`) by copying the default configuration and customizing it. 
- run the app
    ```bash
    ./bin/aurevoir -c path/to/config.toml
    ```

## 📄 License
Distributed under [MIT](https://github.com/leolorenzato/aurevoir/blob/main/LICENSE)