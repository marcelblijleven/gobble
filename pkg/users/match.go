package users

import (
	"errors"
	"log"
)

func MatchUsers(u []*User, m []UserMapping) error {
	if len(u) == 1 {
		return errors.New("only one user found, nothing to match or scrobble")
	}

	mapping := mapFromUserMapping(m)

	for _, i := range u {
		for _, j := range u {
			mappedUsername := mapping[i.Username]

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

func mapFromUserMapping(m []UserMapping) map[string]string {
	mapping := map[string]string{}
	for _, u := range m {
		log.Println(u.String())
		mapping[u.From] = u.To
		mapping[u.To] = u.From
	}

	return mapping
}
