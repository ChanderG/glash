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
  "os/exec"
  "syscall"
)

/*
 * Run outsourced commands.
 * In a seperate function to make it easy to change implementations.
 * INPUT
 * name -- name of command 
 * args -- full string on cli in array form
 */
func outsourceCmd(name string, args []string){
  // using manual exec flow

  // find the binary
  fullname, err := exec.LookPath(name)
  if err != nil {
	fmt.Println("glash: ", name, " command not found")
	os.Exit(0)
  }

  env := os.Environ()

  err = syscall.Exec(fullname, args, env)
  if err != nil {
	panic(err)
  }
}

/*
 * Process commands.
 * INPUT
 * string representing user input
 */
func processCommand(command string) {
  args := strings.Split(command, " ")
  name := args[0]

  // manual processing of in-built commands
  if name == "x" {
	// exit command
	os.Exit(0)
  }

  // outsource commands
  outsourceCmd(name, args)
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
