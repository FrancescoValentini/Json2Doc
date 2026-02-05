package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
)

// renderTemplate processes the template with JSON data
func renderTemplate(tmplContent []byte, jsonData []byte, logger *AppLogger) (string, error) {
	logger.Info("Parsing JSON data")

	// Parse JSON into a generic map
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	logger.Info("JSON parsed successfully (%d top-level keys)", len(data))
	if logger.verbose {
		for key := range data {
			logger.Debug("  - JSON key: %s", key)
		}
	}

	logger.Info("Parsing template")

	// Parse template
	tmpl, err := template.New("main").Parse(string(tmplContent))
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	logger.Info("Template parsed successfully, Executing template")

	// Execute template
	var output bytes.Buffer
	if err := tmpl.Execute(&output, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	result := output.String()
	logger.Info("Template executed successfully (output size: %d bytes)", len(result))

	return result, nil
}
