package main

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
  pattern, line := "^.*line.*$", "this line here"
  if !isMatch(pattern, line) {
    t.Errorf("Expected pattern '%s' to match line '%s', but it did not", pattern, line)
  }
}

func TestIsMatchNoMatch(t *testing.T) {
  pattern, line := "^.*statement.*$", "this line here"
  if isMatch(pattern, line) {
    t.Errorf("Expected pattern '%s' to not match line '%s', but it did", pattern, line)
  }
}

func TestIsNotMatch(t *testing.T) {
  pattern, line := "^.*line.*$", "this line here"
  if isNotMatch(pattern, line) {
    t.Errorf("Expected pattern '%s' to match line '%s', but it did not", pattern, line)
  }
}

func TestIsNotMatchWithMatch(t *testing.T) {
  pattern, line := "^.*statement.*$", "this line here"
  if !isNotMatch(pattern, line) {
    t.Errorf("Expected pattern '%s' to not match line '%s', but it did", pattern, line)
  }
}

func TestIsMatchLines(t *testing.T){
  pattern, lines := "^.*line.*$", []string{"one line", "two line", "three three"}
  matches := filterLines(pattern, lines, isMatch)
  if len(matches) != 2 {
    t.Errorf("Expected 2 matches, but got %d", len(matches))
  }
}

func TestIsNotMatchLines(t *testing.T){
  pattern, lines := "^.*line.*$", []string{"one line", "two line", "three three"}
  matches := filterLines(pattern, lines, isNotMatch)
  if len(matches) != 1 {
    t.Errorf("Expected 1 matches, but got %d", len(matches))
  }
}

func TestArgumentsErrorFormatting(t *testing.T) {
	err := ArgumentsError{"Testing"}
	output := err.Error()
	if output != "Testing" {
		t.Errorf("Expected 'Testing', got '%s'", output)
	}
}

func TestArgumentsProcessing(t *testing.T) {
	path, pattern := "./gogrep.go", "func"
	arguments, err := processArguments([]string{pattern, path})

	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if arguments.Path != path {
		t.Errorf("Expected '%s', got %s", path, arguments.Path)
	}

	if arguments.Pattern != pattern {
		t.Errorf("Expected '%s', got %s", pattern, arguments.Pattern)
	}

  if arguments.Inverse {
		t.Errorf("Expected Inverse to not be set, but it was set")
  }
}

func TestGetMatchingLinesMany(t *testing.T) {
	lines := []string{"first line", "second", "third line"}
	pattern := "line"
	expectedLines := []string{"first line", "third line"}

	matchingLines := getMatchingLines(pattern, lines)

  if len(matchingLines) != len(expectedLines) {
    t.Errorf("Expected %d lines, got %d", len(expectedLines), len(matchingLines))
  }
}

func TestGetMatchingLinesNoLines(t *testing.T) {
	lines := []string{}
	pattern := "line"
	expectedLines := []string{}

	matchingLines := getMatchingLines(pattern, lines)

  if len(matchingLines) != len(expectedLines) {
    t.Errorf("Expected %d lines, got %d", len(expectedLines), len(matchingLines))
	}
}

func TestGetMatchingLinesOneLine(t *testing.T) {
	lines := []string{"only line"}
	pattern := "line"
	expectedLines := []string{"only line"}

	matchingLines := getMatchingLines(pattern, lines)

  if len(matchingLines) != len(expectedLines) {
    t.Errorf("Expected %d lines, got %d", len(expectedLines), len(matchingLines))
  }
}

func TestGetMatchingLinesMatchOneLine(t *testing.T) {
	lines := []string{"one line", "two line", "red line", "blue line"}
	pattern := "red"
	expectedLines := []string{"red line"}

	matchingLines := getMatchingLines(pattern, lines)

  if len(matchingLines) != len(expectedLines) {
    t.Errorf("Expected %d lines, got %d", len(expectedLines), len(matchingLines))
  }
}
