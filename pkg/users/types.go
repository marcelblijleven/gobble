package users

import (
	"fmt"
)

// User is a Gobble user, can be from any map
// mapped users are also stored on the User object
type User struct {
	ID           string        `json:"id"`
	Username     string        `json:"username"`
	MatchedUsers []MatchedUser `json:"matched_users"`
	*Source
}

// GetIdentifier combines the users' Username and Source.ServerID to create a unique
// identifier per app
func (u User) GetIdentifier() string {
	return fmt.Sprintf("%s:%s", u.Username, u.ServerID)
}

func (u User) String() string {
	return fmt.Sprintf("user %q from service %s with server id %q", u.Username, u.ServerType, u.ServerID)
}

// MatchedUser defines the match between two User objects
type MatchedUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	*Source
}

// Source defines the fields for the User source, like server id and server type
type Source struct {
	ServerID   string `json:"server_id"`
	ServerType string `json:"server_type"`
}
