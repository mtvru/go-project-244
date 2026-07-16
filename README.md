# Difference Calculator (gendiff)

A CLI utility written in Go that compares two configuration files (JSON or YAML)
and shows the differences between them in various output formats.

### Hexlet tests and linter status:
[![Actions Status](https://github.com/mtvru/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/mtvru/go-project-244/actions)
[![main](https://github.com/mtvru/go-project-244/actions/workflows/main.yml/badge.svg)](https://github.com/mtvru/go-project-244/actions/workflows/main.yml)

## Features

- Supported input formats: **JSON** and **YAML** (`.yml`, `.yaml`)
- Report output formats: **stylish** (default), **plain**, **json**
- Works with both relative and absolute paths
- Compares flat and nested data structures

## Installation

```bash
make install
make build
```

The executable will be built at `bin/gendiff`.

## Usage

```bash
# default format (stylish)
./bin/gendiff filepath1.json filepath2.json

# explicit format selection
./bin/gendiff --format plain filepath1.yml filepath2.yml
./bin/gendiff --format json filepath1.json filepath2.json

# help
./bin/gendiff --help
```
