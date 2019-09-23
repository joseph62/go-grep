package main

import (
  "fmt"
  "os"
)

type ArgumentsError struct {
  Arguments []string
  Message string
}

func (e *ArgumentsError) Error() string {
  return fmt.Sprintf("Error: %s", e.Message)
}

type Arguments struct {
  Path string
  Pattern string
}

func printError(e error, help bool){
  fmt.Fprintf(os.Stderr, "%s\n", e.Error())
  if (help){
    printHelp()
  }
}

func printHelp() {
  fmt.Fprintf(os.Stderr, "Usage: gogrep <file> <pattern>\n")
}

func processArguments(args []string) (Arguments, error) {
  if len(args) != 3 {
    return Arguments{}, &ArgumentsError{args, "Too many or too few arguments"}
  }
  return Arguments{args[1], args[2]}, nil
}

func main() {
  arguments, err := processArguments(os.Args)
  if err != nil {
    printError(err, true)
    return
  }
  fmt.Printf("Path: %s, Pattern: %s\n", arguments.Path, arguments.Pattern)
}
