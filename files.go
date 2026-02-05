package main

import (
	"fmt"
	"os"
)

// readFile reads a file and returns its content with logging
func readFile(filepath string, logger *AppLogger) ([]byte, error) {
	logger.Info("Reading file: %s", filepath)

	info, err := os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist: %s", filepath)
		}
		return nil, fmt.Errorf("failed to access file %s: %w", filepath, err)
	}

	if info.IsDir() {
		return nil, fmt.Errorf("path is a directory, not a file: %s", filepath)
	}

	logger.Info("File size: %d bytes", info.Size())

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filepath, err)
	}

	logger.Info("Successfully read file: %s (%d bytes)", filepath, len(data))
	return data, nil
}

// writeFile writes data to a file with logging
func writeFile(filepath string, data []byte, logger *AppLogger) error {
	logger.Info("Writing to file: %s", filepath)

	// Check if output directory exists
	dir := filepath[:len(filepath)-len(filepath[len(filepath)-1:])]
	if dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filepath, err)
	}

	logger.Info("Successfully wrote output file: %s (%d bytes)", filepath, len(data))
	return nil
}
