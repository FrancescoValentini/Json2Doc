package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	cfg := parseFlags() //Parses command-line flags and load the configuration

	// Show version and exit if requested
	if cfg.ShowVersion {
		fmt.Fprintf(os.Stderr, "%s v%s by %s\n\n", APP_NAME, APP_VERSION, APP_AUTHOR)
		os.Exit(0)
	}

	logger := NewLogger(cfg.LogDebug) // Instantiate the logger

	startTime := time.Now()
	logger.Info("Starting %s v%s", APP_NAME, APP_VERSION)

	// TODO: Main

	endTime := time.Since(startTime)
	logger.Info("Elapsed Time: %s", endTime)
	os.Exit(0)
}

// parseFlags parses command-line flags
func parseFlags() *AppConfig {
	cfg := &AppConfig{}

	flag.StringVar(&cfg.InputFile, "i", "", "Input JSON file (required)")
	flag.StringVar(&cfg.TemplateFile, "t", "", "Template file (required)")
	flag.StringVar(&cfg.OutputFile, "o", "", "Output file (required)")
	flag.BoolVar(&cfg.LogDebug, "d", false, "Enable verbose logging (debug)")
	flag.BoolVar(&cfg.ShowVersion, "v", false, "Show version information")

	// -h flag View all available flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s v%s by %s\n\n", APP_NAME, APP_VERSION, APP_AUTHOR)
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "Renders a template file using data from a JSON file.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -i data.json -t template.tpl -o output.txt\n", filepath.Base(os.Args[0]))
	}

	flag.Parse()
	return cfg
}
