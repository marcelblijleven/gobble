package watcher

import (
	"fmt"
	"gobble/pkg/tasks"
	"gobble/pkg/users"
	"log"
)

func (w *Watcher) initializeUsers() error {
	log.Println("initializing users")

	u, err := tasks.GetUsers(w.RegisteredServices)

	if err != nil {
		fmt.Println("ERREREREREREREr", err)
		return err
	}

	if err = users.MatchUsers(u, w.Config.UserMappings); err != nil {
		fmt.Println("ERREREREREREREr", err)
		return err
	}
	fmt.Printf("MIAUW %+v\n", u)
	w.Users = u

	return nil
}
