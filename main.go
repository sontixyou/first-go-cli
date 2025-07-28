package main

import (
	"fmt"
	"os"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Handler     func(args []string) error
}

type CLI struct {
	Name        string
	Description string
	Commands    map[string]*Command
}

func NewCLI(name, description string) *CLI {
	return &CLI{
		Name:        name,
		Description: description,
		Commands:    make(map[string]*Command),
	}
}

func (cli *CLI) AddCommand(name, description string, handler func(args []string) error) {
	cli.Commands[name] = &Command{
		Name:        name,
		Description: description,
		Handler:     handler,
	}
}

func (cli *CLI) Run() {
	args := os.Args[1:]
	
	if len(args) == 0 {
		cli.showHelp()
		return
	}

	commandName := args[0]
	commandArgs := args[1:]

	if commandName == "help" || commandName == "-h" || commandName == "--help" {
		cli.showHelp()
		return
	}

	command, exists := cli.Commands[commandName]
	if !exists {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", commandName)
		fmt.Fprintf(os.Stderr, "Run '%s help' for usage.\n", cli.Name)
		os.Exit(1)
	}

	if err := command.Handler(commandArgs); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func (cli *CLI) showHelp() {
	fmt.Printf("%s - %s\n\n", cli.Name, cli.Description)
	fmt.Println("Usage:")
	fmt.Printf("  %s <command> [arguments]\n\n", cli.Name)
	fmt.Println("Available commands:")
	
	for name, cmd := range cli.Commands {
		fmt.Printf("  %-12s %s\n", name, cmd.Description)
	}
	
	fmt.Printf("  %-12s %s\n", "help", "Show this help message")
}

func main() {
	cli := NewCLI("mycli", "A simple CLI tool built with Go standard library")

	cli.AddCommand("hello", "Print a greeting message", func(args []string) error {
		name := "World"
		if len(args) > 0 {
			name = strings.Join(args, " ")
		}
		fmt.Printf("Hello, %s!\n", name)
		return nil
	})

	cli.AddCommand("version", "Show version information", func(args []string) error {
		fmt.Println("mycli version 1.0.0")
		return nil
	})

	cli.AddCommand("echo", "Echo the provided arguments", func(args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no arguments provided")
		}
		fmt.Println(strings.Join(args, " "))
		return nil
	})

	cli.Run()
}