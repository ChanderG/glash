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
)

/*
 * handler for command "cd"
 * INPUT
 * standard format of name and args
 * KNOWN ISSUES
 * Does not handle - ie : "cd -" fails
 */
func handleCd(name string, args []string){
  dest := ""
  if len(args) == 1 {
	// in case, no extra args 
	// simulate cd to home dir
	var exists bool 
    dest, exists = os.LookupEnv("HOME")
	if !exists {
	  fmt.Println("home dir not found")
	  return
	}
  } else {
	dest = args[1]
  }

  err := os.Chdir(dest)
  if err != nil {
	fmt.Println(err)
  }
}

/*
 * Run outsourced commands.
 * In a seperate function to make it easy to change implementations.
 * INPUT
 * name -- name of command 
 * args -- full string on cli in array form
 */
func outsourceCmd(name string, args []string){

  cmd := exec.Command(name)
  cmd.Args = args

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  err := cmd.Run()
  if err != nil {
	fmt.Println(err)
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

  // process builtins follwed by outsourcing
  switch name {
  case "x":  // exit command
	  os.Exit(0)
  case "c":  // clear command
	  outsourceCmd("clear", []string{"clear"})
  case "cd": 
	  handleCd(name, args)
  default:   // outsource
      outsourceCmd(name, args)
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
