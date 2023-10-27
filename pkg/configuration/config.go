package configuration

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"gobble/pkg/users"
)

// Config is a representation of the Gobble config file
type Config struct {
	Host string `json:"host" toml:"host"`
	Port string `json:"port" toml:"port"`
	*Services
	*users.UserConfig
}

// AdditionalConfig stores service/integration agnostic configuration
type AdditionalConfig struct {
	Name     string `json:"name" toml:"name"`
	Timeout  int    `json:"timeout" toml:"duration"`
	Interval int    `json:"interval" toml:"interval"`
}

// New creates a new Config with defaults
func New() *Config {
	return &Config{
		Host: "",
		Port: "",
		Services: &Services{
			Jellyfin: []*JellyfinConfig{},
			Plex:     []*PlexConfig{},
		},
		UserConfig: &users.UserConfig{UserMappings: []users.UserMapping{}},
	}
}

// Parse reads the config file
func (c *Config) Parse(flag *Flags) error {
	if flag.ConfigFile == "" {
		return errors.New("no configuration file found")
	}

	_, err := toml.DecodeFile(flag.ConfigFile, c)

	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err = c.Setup(); err != nil {
		return err
	}

	return nil
}
