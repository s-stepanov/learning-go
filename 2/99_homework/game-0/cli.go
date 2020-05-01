package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CLI handles IO to stdin/out
type CLI struct {
	reader *bufio.Reader
}

// NewCLI constructor
func NewCLI () *CLI {
	cli := new(CLI)
	cli.reader = bufio.NewReader(os.Stdin)

	return cli
}

// ParseUserInput Parses user's game command from standart input
func (cli *CLI) ParseUserInput() (command string, parameters []string) {
	input, _ := cli.reader.ReadString('\n')
	commands := strings.Split(input, " ")
	fmt.Println(commands)

	return commands[0], commands[1:]
}

// GetUserNickname gets user's nickname on game startup
func (cli *CLI) GetUserNickname() string {
	fmt.Println("Введите имя персонажа")
	nickname, _ := cli.reader.ReadString('\n')
	return nickname
}