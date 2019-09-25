package main

import (
  "testing"
)

func TestArgumentsErrorFormatting(t *testing.T){
  err := ArgumentsError{"Testing"}
  output := err.Error()
  if output != "Error: Testing" {
    t.Errorf("Expected 'Error: Testing', got '%s'", output)
  }
}

func TestArgumentsProcessing(t *testing.T){
  path, pattern := "./gogrep.go", "third"
  arguments, err := processArguments([]string{"first", path, pattern})

  if err != nil {
    t.Errorf("Expected no error, got %s", err.Error())
  }

  if arguments.Path != path {
    t.Errorf("Expected '%s', got %s", path, arguments.Path)
  }

  if arguments.Pattern != pattern {
    t.Errorf("Expected '%s', got %s", pattern, arguments.Pattern)
  }
}

func TestGetMatchingLinesMany(t *testing.T){
  lines := []string{"first line", "second", "third line"}
  pattern := "line"
  expectedLines := []string{"first line", "third line"}

  matchingLines, err := getMatchingLines(pattern, lines)

  if err != nil {
    t.Errorf("Expected no error, got %s", err.Error())
  } else {
    if len(matchingLines) != len(expectedLines) {
      t.Errorf("Expected %d lines, got %d", len(expectedLines), len(matchingLines))
    }
  }
}

func TestGetMatchingLinesNoLines(t *testing.T){
  lines := []string{}
  pattern := "line"
  expectedLines := []string{}

  matchingLines, err := getMatchingLines(pattern, lines)

  if err != nil {
    t.Errorf("Expected no error, got %s", err.Error())
  } else {
    if len(matchingLines) != len(expectedLines) {
      t.Errorf("Expected %d lines, got %d", len(expectedLines), len(matchingLines))
    }
  }
}

func TestGetMatchingLinesOneLine(t *testing.T){
  lines := []string{"only line"}
  pattern := "line"
  expectedLines := []string{"only line"}

  matchingLines, err := getMatchingLines(pattern, lines)

  if err != nil {
    t.Errorf("Expected no error, got %s", err.Error())
  } else {
    if len(matchingLines) != len(expectedLines) {
      t.Errorf("Expected %d lines, got %d", len(expectedLines), len(matchingLines))
    }
  }
}

func TestGetMatchingLinesMatchOneLine(t *testing.T){
  lines := []string{"one line", "two line", "red line", "blue line"}
  pattern := "red"
  expectedLines := []string{"red line"}

  matchingLines, err := getMatchingLines(pattern, lines)

  if err != nil {
    t.Errorf("Expected no error, got %s", err.Error())
  } else {
    if len(matchingLines) != len(expectedLines) {
      t.Errorf("Expected %d lines, got %d", len(expectedLines), len(matchingLines))
    }
  }
}
