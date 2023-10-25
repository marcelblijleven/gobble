package jellyfin

import "gobble/pkg/users"

// convertUsers takes Jellyfin user objects and converts them
// to the gobble User object
func convertUsers(data []User) []users.User {
	var converted []users.User

	for _, u := range data {
		converted = append(converted, users.User{
			ID:       u.Id,
			Username: u.Name,
		})
	}

	return converted
}
