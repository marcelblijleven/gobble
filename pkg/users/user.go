package users

// User is a Gobble user, can be from any map
// mapped users are also stored on the User object
type User struct {
	ID           string        `json:"id"`
	Username     string        `json:"username"`
	MatchedUsers []MatchedUser `json:"matched_users"`
	*Source
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
