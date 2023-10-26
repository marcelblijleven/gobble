package main

import (
	"gobble/pkg/watcher"
	"log"
)

func main() {
	if err := watcher.Run(); err != nil {
		log.Fatalln(err)
	}
}
