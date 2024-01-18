package config

// KeyNames is a struct that contains the names of keys.
type KeyNames struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName    string
	ApplicationWebsite string
	LogFilePath        string
	SoftwareVersion    string
}

// Keys contains the names of config keys.
var Keys = KeyNames{
	ConfigPath: "config-path", // CLI only
	LogLevel:   "log-level",

	// application
	ApplicationName:    "application-name",
	ApplicationWebsite: "application-website",
	LogFilePath:        "log-file-path",
	SoftwareVersion:    "software-version", // Set at build
}
