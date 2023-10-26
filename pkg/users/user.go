package users

// User is a Gobble user, can be from any map
// mapped users are also stored on the User object
type User struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Source      string `json:"source"`
	MappedUsers []User `json:"mapped_users"`
}
