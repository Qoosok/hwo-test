package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"testing"
)

const expectedOutput = "info: 5\nwarning: 3\nerror: 1\n"

func TestParseFlags(t *testing.T) {
	os.Setenv("LOG_ANALYZER_FILE", "test.log")
	os.Setenv("LOG_ANALYZER_LEVEL", "error")
	os.Setenv("LOG_ANALYZER_OUTPUT", "output.txt")
	defer os.Unsetenv("LOG_ANALYZER_FILE")
	defer os.Unsetenv("LOG_ANALYZER_LEVEL")
	defer os.Unsetenv("LOG_ANALYZER_OUTPUT")

	flag.CommandLine = flag.NewFlagSet("test", flag.ExitOnError)
	filePath, logLevel, outputPath, err := parseFlags()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if filePath != "test.log" {
		t.Errorf("expected filePath to be 'test.log', got %s", filePath)
	}
	if logLevel != "error" {
		t.Errorf("expected logLevel to be 'error', got %s", logLevel)
	}
	if outputPath != "output.txt" {
		t.Errorf("expected outputPath to be 'output.txt', got %s", outputPath)
	}
}

func TestProcessLogFile(t *testing.T) {
	content := `info: all good
error: something went wrong
error: another error`
	err := os.WriteFile("test.log", []byte(content), 0o644)
	if err != nil {
		t.Fatalf("could not create test file: %v", err)
	}
	defer os.Remove("test.log")

	levelCounts, err := processLogFile("test.log", "error")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if levelCounts["error"] != 2 {
		t.Errorf("expected 2 errors, got %d", levelCounts["error"])
	}
}

func TestGenerateStatistics(t *testing.T) {
	levelCounts := map[string]int{"info": 5, "warning": 3, "error": 1}
	result := generateStatistics(levelCounts)
	if result != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, result)
	}
}

func TestOutputResultsToFile(t *testing.T) {
	err := outputResults(expectedOutput, "test_output.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	data, err := os.ReadFile("test_output.txt")
	if err != nil {
		t.Fatalf("could not read output file: %v", err)
	}
	if string(data) != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, data)
	}
	os.Remove("test_output.txt")
}

func TestOutputResultsToStdout(t *testing.T) {
	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w

	// Run the outputResults function in a goroutine to write to stdout
	go func() {
		defer w.Close() // Ensure the pipe is closed after writing
		outputResults(expectedOutput, "")
	}()

	// Read the output from the pipe
	var buf strings.Builder
	_, err := io.Copy(&buf, r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Restore the original stdout
	os.Stdout = originalStdout

	// Compare the captured output
	if buf.String() != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, buf.String())
	}
}
