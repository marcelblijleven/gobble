package tasks

import (
	"fmt"
	"gobble/pkg/apps"
	"gobble/pkg/configuration"
	"gobble/pkg/users"
)

// GetUsers calls the GetUser method on all the provided apps,
// combines the result into one slice of users.User and returns it
func GetUsers(cfg *configuration.Config) ([]users.User, error) {
	var u []users.User

	// TODO: programmatically retrieve apps instead of manually defining slice of Apps
	for _, s := range []apps.App{cfg.Jellyfin.App, cfg.Plex.App} {
		su, err := s.GetUsers()

		if err != nil {
			return nil, err
		}

		u = append(u, su...)
	}

	fmt.Printf("Foo %+v\n", u)
	return u, nil
}
