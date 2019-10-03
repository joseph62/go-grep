package main

import (
	"fmt"
  "flag"
	"io/ioutil"
  "io"
	"os"
	"regexp"
	"strings"
)

func gogrepUsage(defaultUsage func(), writer io.Writer) func() {
  return func() {
    fmt.Fprintf(writer, "gogrep pattern path [--inverse]\n")
    defaultUsage()
  }
}

type LineMatch struct {
	Line       string
	LineNumber uint64
}

type ArgumentsError struct {
	Message string
}

func (e *ArgumentsError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

type Arguments struct {
	Pattern string
	Path    string
  Inverse bool
}

func processArguments(args []string) (Arguments, error) {
  parser := flag.NewFlagSet("gogrep", flag.ExitOnError)
  parser.Usage = gogrepUsage(parser.Usage, os.Stderr)
  
  inverse := parser.Bool("inverse", false, "Print lines that do not match the pattern")

  err := parser.Parse(args)

  if err != nil {
    return Arguments{}, err
  }

  positionalArguments := parser.Args()

	if len(positionalArguments) < 2 {
		return Arguments{}, &ArgumentsError{"Too many or too few arguments"}
	}

	return Arguments{positionalArguments[0], positionalArguments[1], *inverse}, nil
}

func readLines(path string) ([]string, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(contents), "\n"), nil
}

func isMatch(pattern, line string) bool {
  matched, err := regexp.MatchString(pattern, line)
  return err == nil && matched
}

func isNotMatch(pattern, line string) bool {
  return !isMatch(pattern, line)
}

func filterLines(pattern string, lines []string, predicate func(string, string) bool) []string {
  results := []string{}
  for i := 0; i < len(lines); i++ {
    if predicate(pattern, lines[i]) {
      results = append(results, lines[i])
    }
  }
  return results
}

func processMatchingLines(lines []string) []LineMatch {
  results := []LineMatch{}
  for i := 0; i < len(lines); i++ {
    results = append(results, LineMatch{lines[i], uint64(i)})
  }
  return results
}

func getNonMatchingLines(pattern string, lines []string) []LineMatch {
  return processMatchingLines(filterLines(pattern, lines, isNotMatch))
}

func getMatchingLines(pattern string, lines []string) []LineMatch {
  return processMatchingLines(filterLines(pattern, lines, isMatch))
}

func main() {
	arguments, err := processArguments(os.Args)
	if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	lines, err := readLines(arguments.Path)
	if err != nil {
    fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

  var matches []LineMatch
  if arguments.Inverse {
    matches = getNonMatchingLines(arguments.Pattern, lines)
  } else {
    matches = getMatchingLines(arguments.Pattern, lines)
  }

	for i := 0; i < len(matches); i++ {
		match := matches[i]
		fmt.Printf("%d %s\n", match.LineNumber, match.Line)
	}
}
