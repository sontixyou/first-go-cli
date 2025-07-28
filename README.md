# mycli

A simple CLI tool built with Go standard library

## Overview

This is a basic command-line interface (CLI) application written in Go using only the standard library. The CLI provides a framework for creating and managing multiple commands with their own handlers.

## Features

- Command-based interface with help system
- Built-in help command (`help`, `-h`, `--help`)
- Error handling and user-friendly error messages
- Extensible command system

## Available Commands

- `hello [name]` - Print a greeting message (defaults to "World" if no name provided)
- `version` - Show version information (currently v1.0.0)
- `echo <text>` - Echo the provided arguments
- `help` - Show help message with all available commands

## Usage

```bash
# Show help
go run main.go help

# Say hello
go run main.go hello
go run main.go hello Solid Snake

# Show version
go run main.go version

# Echo text
go run main.go echo This is a test message
```

## Building

```bash
go build -o mycli main.go
```

## Running

```bash
./mycli <command> [arguments]
```

## Code Structure

The application consists of:

- `Command` struct: Represents individual commands with name, description, and handler
- `CLI` struct: Main CLI application with command management
- Command registration system for easy extensibility
- Built-in help system and error handling

## License

This project is a simple demonstration CLI tool.

