package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func printError(e error, help bool) {
	fmt.Fprintf(os.Stderr, "%s\n", e.Error())
	if help {
		printHelp()
	}
}

func printHelp() {
	fmt.Fprintf(os.Stderr, "Usage: gogrep <file> <pattern>\n")
}

type LineMatch struct {
	Line       string
	LineNumber uint64
}

type ArgumentsError struct {
	Message string
}

func (e *ArgumentsError) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

type Arguments struct {
	Path    string
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
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}

func getMatchingLines(pattern string, lines []string) ([]LineMatch, error) {
	matches := []LineMatch{}
	for i := uint64(0); i < uint64(len(lines)); i++ {
		if matched, err := regexp.MatchString(pattern, lines[i]); err != nil {
			return matches, err
		} else if matched {
			matches = append(matches, LineMatch{lines[i], i})
		}
	}
	return matches, nil
}

func main() {

	arguments, err := processArguments(os.Args)
	if err != nil {
		printError(err, true)
		return
	}

	lines, err := readLines(arguments.Path)
	if err != nil {
		printError(err, false)
		return
	}

	matches, err := getMatchingLines(arguments.Pattern, lines)
	if err != nil {
		printError(err, false)
		return
	}

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		fmt.Printf("%d %s\n", match.LineNumber, match.Line)
	}
}
