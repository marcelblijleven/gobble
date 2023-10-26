package common

import "fmt"

type SystemInfo struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Version         string `json:"version"`
	OperatingSystem string `json:"operating_system"`
}

func (i *SystemInfo) String() string {
	return fmt.Sprintf("name: %s (version: %s, os: %s)", i.Name, i.Version, i.OperatingSystem)
}
