// hw12_fixing
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func parseFlags() (filePath, logLevel, outputPath string, err error) {
	filePath = os.Getenv("LOG_ANALYZER_FILE")
	logLevel = os.Getenv("LOG_ANALYZER_LEVEL")
	outputPath = os.Getenv("LOG_ANALYZER_OUTPUT")

	flag.StringVar(&filePath, "file", filePath, "Path to the log file")
	flag.StringVar(&logLevel, "level", logLevel, "Log level to analyze (default: 'info')")
	flag.StringVar(&outputPath, "output", outputPath, "Path to the output file")

	flag.Parse()

	if filePath == "" {
		return "", "", "", fmt.Errorf(
			"log file path must be specified either via -file flag or LOG_ANALYZER_FILE environment variable")
	}

	if logLevel == "" {
		logLevel = "info"
	}

	return filePath, logLevel, outputPath, nil
}

func processLogFile(filePath, logLevel string) (map[string]int, error) {
	levelCounts := map[string]int{
		"info":    0,
		"warning": 0,
		"error":   0,
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, logLevel) {
			levelCounts[logLevel]++
		}
	}

	scannerErr := scanner.Err()
	if scannerErr != nil {
		return nil, scannerErr
	}

	return levelCounts, nil
}

func generateStatistics(levelCounts map[string]int) string {
	var result strings.Builder
	for level, count := range levelCounts {
		fmt.Fprintf(&result, "%s: %d\n", level, count)
	}
	return result.String()
}

func outputResults(statistics, outputPath string) error {
	if outputPath != "" {
		file, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString(statistics)
		if err != nil {
			return err
		}
	} else {
		fmt.Print(statistics)
	}
	return nil
}

func main() {
	filePath, logLevel, outputPath, err := parseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	levelCounts, err := processLogFile(filePath, logLevel)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	statistics := generateStatistics(levelCounts)

	err = outputResults(statistics, outputPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
