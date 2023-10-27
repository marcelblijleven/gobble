package users

import "fmt"

// UserConfig represents user related config file entries
type UserConfig struct {
	UserMappings map[string]string `json:"user_mapping" toml:"user_mapping"`
}

// UserMapping is part of the UserConfig and allows usernames to be mapped between apps
type UserMapping struct {
	From string `json:"from" toml:"from"`
	To   string `json:"to" toml:"to"`
}

func (u UserMapping) String() string {
	return fmt.Sprintf("username %q mapped to %q, and vice versa", u.From, u.To)
}
