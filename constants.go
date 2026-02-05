package main

const (
	// Application Details
	APP_VERSION = "1.0.0"
	APP_NAME    = "Json2Doc"
	APP_AUTHOR  = "Francesco Valentini"

	// Logger levels
	LOG_LEVEL_DEBUG = "DEBUG"
	LOG_LEVEL_INFO  = "INFO"
	LOG_LEVEL_WARN  = "WARN"
	LOG_LEVEL_ERROR = "ERROR"
)

// AppConfig holds application configuration
type AppConfig struct {
	InputFile    string // Input json file path
	TemplateFile string // Path to the file containing the template
	OutputFile   string // Output file path
	LogDebug     bool   // true if the logger is in debug mode
	ShowVersion  bool   // The application only prints the version
}
