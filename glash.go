/*
 * glash: go lightweight anonymous shell
 * Design objectives:
 * 1. Just basic functionality
 * 2. Heavily customized for my use. Not configuration.
 * 3. No history saved anywhere.
 * 4. Meant to be carried around -- not installed
 * 5. First go project. Have fun.
 */
package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

/*
 * Process commands.
 * INPUT
 * string representing user input
 */
func processCommand(command string) {
  args := strings.Split(command, " ")
  cmd := args[0]

  // manual processing of in-built commands
  if cmd == "x" {
	// exit command
	os.Exit(0)
  }
}

/*
 * Display prompt and obtain input. 
 * INPUT
 * pointer to a Buffered Reader meant to read from stdio
 */
func prompt(conreader *bufio.Reader) {
  // display prompt
  fmt.Print("$ ")

  // obtain input
  input, err := conreader.ReadString('\n')
  if err != nil {
	panic(err)
  }

  // remove trailing '\n'
  input = strings.TrimRight(input, "\n")

  // process command
  processCommand(input)
}

/*
 * Main runner.
 */
func main() {
  fmt.Println("> Welcome to glash.")

  // reader to read from console
  conreader := bufio.NewReader(os.Stdin)
  for true {
	prompt(conreader)
  }
}
