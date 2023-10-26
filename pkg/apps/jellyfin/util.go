package jellyfin

import (
	"gobble/pkg/common"
	"gobble/pkg/users"
)

// convertUsers takes Jellyfin user objects and converts them
// to the gobble User object
func convertUsers(data []User) []users.User {
	var converted []users.User

	for _, u := range data {
		converted = append(converted, users.User{
			ID:          u.Id,
			Username:    u.Name,
			Source:      "jellyfin",
			MappedUsers: []users.User{},
		})
	}

	return converted
}

// convertSystemInfo takes Jellyfin system info response and
// converts it to a common SystemInfo object
func convertSystemInfo(data SystemInfo) *common.SystemInfo {
	return &common.SystemInfo{
		Id:              data.Id,
		Name:            data.ServerName,
		Version:         data.Version,
		ProductName:     data.ProductName,
		OperatingSystem: data.OperatingSystem,
	}
}
