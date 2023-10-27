package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetIdentifier(t *testing.T) {
	u := &User{
		ID:           "test-id",
		Username:     "test-user",
		MatchedUsers: []MatchedUser{},
		Source: &Source{
			ServerID:   "server-id",
			ServerType: "plex",
		},
	}

	assert.Equal(t, "test-user:server-id", u.GetIdentifier())
}

func TestUser_String(t *testing.T) {

	u := &User{
		ID:           "test-id",
		Username:     "test-user",
		MatchedUsers: []MatchedUser{},
		Source: &Source{
			ServerID:   "server-id",
			ServerType: "plex",
		},
	}

	assert.Equal(
		t,
		"user \"test-user\" from service plex with server id \"server-id\"",
		u.String())
}
