package config

// Config is the plugin configuration.
type Config struct {
	// Headers is a map of headers to add to the request to the upstream.
	Headers map[string]string `json:"headers,omitempty"`
}

// New creates the default plugin configuration.
func New() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}
