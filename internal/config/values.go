package config

// Values contains the type of each value.
type Values struct {
	LogLevel   string
	ConfigPath string

	// application
	ApplicationName    string
	ApplicationWebsite string
	SoftwareVersion    string
}

// Defaults contains the default values.
var Defaults = Values{
	ConfigPath: "",
	LogLevel:   "info",

	// application
	ApplicationName:    "McUberDisc",
	ApplicationWebsite: "https://tyr.codes/tyr/mcuberdisc",
	SoftwareVersion:    "dev",
}
