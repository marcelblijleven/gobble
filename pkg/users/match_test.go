package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchUsers_OneUser(t *testing.T) {
	u := &User{
		ID:           "foo",
		Username:     "bar",
		MatchedUsers: nil,
		Source: &Source{
			ServerID:   "",
			ServerType: "",
		},
	}

	err := MatchUsers([]*User{u}, nil)

	assert.Error(t, err)
	assert.Nil(t, u.MatchedUsers)
}

func TestMatchUsers_DuplicateUsers(t *testing.T) {
	u := &User{
		ID:           "foo",
		Username:     "bar",
		MatchedUsers: []MatchedUser{},
		Source: &Source{
			ServerID:   "server-id",
			ServerType: "plex",
		},
	}
	u2 := &User{
		ID:           "foo",
		Username:     "bar",
		MatchedUsers: []MatchedUser{},
		Source: &Source{
			ServerID:   "server-id",
			ServerType: "plex",
		},
	}

	err := MatchUsers([]*User{u, u2}, nil)

	assert.Nil(t, err)
	assert.Lenf(t, u.MatchedUsers, 0, "expected length to be 0")
	assert.Lenf(t, u2.MatchedUsers, 0, "expected length to be 0")
}

func TestMatchUsers_MatchingUsers(t *testing.T) {
	u := &User{
		ID:           "foo",
		Username:     "bar",
		MatchedUsers: []MatchedUser{},
		Source: &Source{
			ServerID:   "server-id",
			ServerType: "plex",
		},
	}
	u2 := &User{
		ID:           "foo",
		Username:     "bar",
		MatchedUsers: []MatchedUser{},
		Source: &Source{
			ServerID:   "server-id-2",
			ServerType: "plex",
		},
	}

	err := MatchUsers([]*User{u, u2}, nil)

	assert.Nil(t, err)
	assert.Lenf(t, u.MatchedUsers, 1, "expected length to be 1")
	assert.Lenf(t, u2.MatchedUsers, 1, "expected length to be 1")

	assert.Equal(t, u2.Username, u.MatchedUsers[0].Username)
	assert.Equal(t, u2.ID, u.MatchedUsers[0].ID)
	assert.Equal(t, u2.Source, u.MatchedUsers[0].Source)

	assert.Equal(t, u.Username, u2.MatchedUsers[0].Username)
	assert.Equal(t, u.ID, u2.MatchedUsers[0].ID)
	assert.Equal(t, u.Source, u2.MatchedUsers[0].Source)
}

func TestMatchUsers_MatchingUsers_ByMapping(t *testing.T) {
	u := &User{
		ID:           "foo",
		Username:     "user",
		MatchedUsers: []MatchedUser{},
		Source: &Source{
			ServerID:   "server-id",
			ServerType: "plex",
		},
	}
	u2 := &User{
		ID:           "foo",
		Username:     "testuser",
		MatchedUsers: []MatchedUser{},
		Source: &Source{
			ServerID:   "server-id-2",
			ServerType: "plex",
		},
	}

	mapping := map[string]string{
		u.Username: u2.Username,
	}

	err := MatchUsers([]*User{u, u2}, mapping)

	assert.Nil(t, err)
	assert.Lenf(t, u.MatchedUsers, 1, "expected length to be 1")
	assert.Lenf(t, u2.MatchedUsers, 1, "expected length to be 1")

	assert.Equal(t, u2.Username, u.MatchedUsers[0].Username)
	assert.Equal(t, u2.ID, u.MatchedUsers[0].ID)
	assert.Equal(t, u2.Source, u.MatchedUsers[0].Source)

	assert.Equal(t, u.Username, u2.MatchedUsers[0].Username)
	assert.Equal(t, u.ID, u2.MatchedUsers[0].ID)
	assert.Equal(t, u.Source, u2.MatchedUsers[0].Source)
}
