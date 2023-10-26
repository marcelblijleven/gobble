package configuration

import (
	"gobble/pkg/apps/jellyfin"
	"gobble/pkg/common"
	"log"
	"net/http"
	"time"
)

type AppConfig interface {
	Setup()
	Enabled() bool
}

type JellyfinConfig struct {
	*jellyfin.Config
	*jellyfin.App
	AdditionalConfig
	SystemInfo *common.SystemInfo
}

// Setup configures a Jellyfin app
func (c *JellyfinConfig) Setup() {
	if c.Timeout == 0 {
		c.Timeout = 30
	}

	c.Client = &http.Client{Timeout: time.Duration(c.Timeout) * time.Second}
	c.App = jellyfin.New(c.Config)

	info, err := c.App.GetSystemInfo()

	if err != nil {
		log.Fatalln(err)
	}

	c.SystemInfo = info
	log.Printf("Jellyfin service info retrieved for %q: %s\n", c.Name, info.String())
}

// Enabled determines if the Jellyfin service can be run
func (c *JellyfinConfig) Enabled() bool {
	enabled := true

	defer func() {
		log.Printf("Jellyfin service enabled: %t\n", enabled)
	}()

	if c == nil {
		log.Println("Jellyfin config not set up")
		enabled = false
	}

	if c.URL == "" {
		log.Println("Jellyfin config has no URL set")
		enabled = false
	}

	if c.Token == "" {
		log.Println("Jellyfin config has no Token set")
		enabled = false
	}

	return enabled
}
