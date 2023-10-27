package configuration

import (
	"fmt"
)

type Services struct {
	Jellyfin []*JellyfinConfig
	Plex     []*PlexConfig
}

// Setup calls the Setup method for each of the Services
func (s *Services) Setup() error {
	// Check Jellyfin services
	for _, j := range s.Jellyfin {
		if !j.Enabled() {
			return fmt.Errorf("the Jellyfin service %q with server id is not enabled", j.Name)
		}
		j.Setup()
	}

	for _, p := range s.Plex {
		if !p.Enabled() {
			return fmt.Errorf("the Plex service %q with server id is not enabled", p.Name)
		}
		p.Setup()
	}

	return nil
}
