package configuration

import "errors"

type Services struct {
	Jellyfin *JellyfinConfig
	Plex     *PlexConfig
}

// Setup calls the Setup method for each of the Services
func (s *Services) Setup() error {
	if !s.Jellyfin.Enabled() {
		return errors.New("the Jellyfin service is not enabled")
	}
	s.Jellyfin.Setup()

	if !s.Plex.Enabled() {
		return errors.New("the plex service is not enabled")
	}
	s.Plex.Setup()
	return nil
}
