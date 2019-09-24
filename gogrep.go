package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
)

func printError(e error, help bool){
  fmt.Fprintf(os.Stderr, "%s\n", e.Error())
  if (help){
    printHelp()
  }
}

func printHelp() {
  fmt.Fprintf(os.Stderr, "Usage: gogrep <file> <pattern>\n")
}

type ArgumentsError struct {
  Message string
}

func (e *ArgumentsError) Error() string {
  return fmt.Sprintf("Error: %s", e.Message)
}

type Arguments struct {
  Path string
  Pattern string
}

func validateArguments(args Arguments) error {
  fileInfo, err := os.Stat(args.Path)
  if err != nil {
    return err
  } else if fileInfo.IsDir() {
    return &ArgumentsError{"The given path is a not a regular file"}
  }
  return nil
}

func processArguments(args []string) (Arguments, error) {
  if len(args) != 3 {
    return Arguments{}, &ArgumentsError{"Too many or too few arguments"}
  }

  arguments := Arguments{args[1], args[2]}

  if err := validateArguments(arguments); err != nil {
    return Arguments{}, err
  }

  return arguments, nil
}

func readLines(path string) ([]string, error) {
  contents, err := ioutil.ReadFile(path)
  if err != nil {
    return  nil, err
  }
  return strings.Split(string(contents), "\n"), nil
}

func main() {
  arguments, err := processArguments(os.Args)
  if err != nil {
    printError(err, true)
    return
  }
  fmt.Printf("Path: %s, Pattern: %s\n", arguments.Path, arguments.Pattern)
  lines, err := readLines(arguments.Path)
  if err != nil {
    printError(err, false)
    return
  }

  for i := 0; i < len(lines); i++ {
    fmt.Printf("%d %s\n", i, lines[i])
  }
}
