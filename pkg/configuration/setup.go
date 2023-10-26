package configuration

import "errors"

type Services struct {
	Jellyfin *JellyfinConfig
}

// Setup calls the Setup method for each of the Services
func (s *Services) Setup() error {
	if !s.Jellyfin.Enabled() {
		return errors.New("the Jellyfin service is not enabled")
	}
	s.Jellyfin.Setup()

	return nil
}
