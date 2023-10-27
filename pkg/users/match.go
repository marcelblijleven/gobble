package users

import (
	"errors"
	"log"
)

// MatchUsers takes all users and username mapping and tries to match each username to
// a username of another service. After finding a match, the User.ID, User.Username and Source
// will be added to a MatchedUser struct
func MatchUsers(u []*User, m map[string]string) error {
	if len(u) == 1 {
		return errors.New("only one user found, nothing to match or scrobble")
	}

	for _, i := range u {
		for _, j := range u {
			mappedUsername := m[i.Username]

			if (i.Username == j.Username || mappedUsername != "" && mappedUsername == j.Username) && i.GetIdentifier() != j.GetIdentifier() {
				i.MatchedUsers = append(i.MatchedUsers, MatchedUser{
					ID:       j.ID,
					Username: j.Username,
					Source:   j.Source,
				})
				log.Printf("matched %s to %s\n", i, j)
			}
		}
	}

	return nil
}
