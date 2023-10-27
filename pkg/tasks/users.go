package tasks

import (
	"gobble/pkg/configuration"
	"gobble/pkg/users"
)

// GetUsers calls the GetUser method on all the provided apps,
// combines the result into one slice of users.User and returns it
func GetUsers(appConfigs map[string]configuration.AppConfig) ([]*users.User, error) {
	var u []*users.User

	for _, app := range appConfigs {
		appUsers, err := app.GetApp().GetUsers()

		if err != nil {
			return nil, err
		}

		u = append(u, appUsers...)
	}

	return u, nil
}
