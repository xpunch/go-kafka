package kafka

import (
	"time"

	"github.com/xpunch/configor/toml"
)

// Config represents the kafka configuration.
type Config struct {
	Brokers string        `toml:"brokers"`
	GroupID string        `toml:"groupId"`
	Topics  string        `toml:"topics"`
	Timeout toml.Duration `toml:"timeout"`
}

// NewConfig builds a new configuration with default values.
func NewConfig() *Config {
	return &Config{
		Brokers: "127.0.0.1:9092",
		GroupID: "test-group",
		Topics:  "test-messages",
		Timeout: toml.Duration(1 * time.Minute),
	}
}
