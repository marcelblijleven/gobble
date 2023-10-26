package configuration

import (
	flag "github.com/spf13/pflag"
)

// Flags represents the flags provided to the application
type Flags struct {
	*flag.FlagSet
	ConfigFile string
	DryRun     bool
}
