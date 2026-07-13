package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
)

// renderTemplate processes the template with JSON data
func renderTemplate(tmplContent []byte, jsonData []byte, logger *AppLogger) (string, error) {
	logger.Info("Parsing JSON data")

	// Parse JSON into a generic map
	var data interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	showDataType(data, logger)

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

func showDataType(data interface{}, logger *AppLogger) {
	switch v := data.(type) {

	case map[string]interface{}:
		logger.Info("JSON parsed successfully (%d top-level keys)", len(v))
		if logger.verbose {
			for key := range v {
				logger.Debug("  - JSON key: %s", key)
			}
		}

	case []interface{}:
		logger.Info("JSON parsed successfully (root array, %d elements)", len(v))
		if logger.verbose {
			for i, el := range v {
				logger.Debug("  - element #%d type: %T", i, el)
			}
		}

	default:
		logger.Warn("JSON parsed, root type is %T", v)
	}
}
