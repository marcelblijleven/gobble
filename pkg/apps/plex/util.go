package plex

import (
	"gobble/pkg/users"
	"strconv"
)

// externalUserToUser converts the response from /users to a slice of gobble users
func externalUserToUser(e ExternalUsers) []users.User {
	var u []users.User

	for _, user := range e.User {
		u = append(u, users.User{
			ID:       user.ID,
			Username: user.Username,
			Source: &users.Source{
				ServerID:   e.MachineIdentifier,
				ServerType: "plex",
			},
		})
	}

	return u
}

// plexUserToUser converts the response from /v2/user to a gobble user
func plexUserToUser(u User, serverID string) users.User {
	return users.User{
		ID:       strconv.Itoa(u.Id),
		Username: u.Username,
		Source: &users.Source{
			ServerID:   serverID,
			ServerType: "plex",
		},
	}
}

// combineUsers is a helper method that combines a slice of users with n users from a variadic parameter
func combineUsers(other []users.User, u ...users.User) []users.User {
	return append(u, other...)
}
