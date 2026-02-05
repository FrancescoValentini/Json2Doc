package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filepath, err)
	}

	logger.Info("Successfully read file: %s (%d bytes)", filepath, len(data))
	return data, nil
}

// writeFile writes data to a file with logging
func writeFile(path string, data []byte, logger *AppLogger) error {
	logger.Info("Writing to file: %s", path)

	dir := filepath.Dir(path)

	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
		logger.Info("Created directory: %s", dir)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}

	logger.Info("Successfully wrote output file: %s (%d bytes)", path, len(data))
	return nil
}
