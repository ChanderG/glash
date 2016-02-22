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
  fmt.Println(input)
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