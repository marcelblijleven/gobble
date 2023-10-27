package plex

import (
	"fmt"
	"gobble/pkg/common"
	"gobble/pkg/users"
	"gobble/pkg/util"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// App represents a Plex media app
type App struct {
	config Config
}

// Config represents a config file entry for a Plex app
type Config struct {
	URL    string `json:"url" toml:"url"`
	Token  string `json:"token" toml:"token"`
	Client *http.Client
}

// New creates a new Plex app from the provided config
func New(config *Config) *App {
	if config.Client == nil {
		config.Client = &http.Client{Timeout: 30 * time.Second}
	}

	return &App{config: *config}
}

// GetSystemInfo retrieves system info from the Plex API
func (a *App) GetSystemInfo() (*common.SystemInfo, error) {
	identity, err := a.getServerIdentity()

	if err != nil {
		return nil, err
	}

	caps, err := a.getServerCapabilities()

	if err != nil {
		return nil, err
	}

	return &common.SystemInfo{
		Id:              identity.MediaContainer.MachineIdentifier,
		Name:            caps.MediaContainer.FriendlyName,
		Version:         caps.MediaContainer.Version,
		OperatingSystem: caps.MediaContainer.Platform,
	}, nil
}

// GetUsers retrieves both the 'MyPlex' account user and any other user known to the system
func (a *App) GetUsers() ([]users.User, error) {
	ext, err := a.getUsers()

	if err != nil {
		return nil, err
	}

	me, err := a.getMe()

	if err != nil {
		return nil, err
	}

	return combineUsers(ext, me), nil
}

// getServerIdentity calls the /identity endpoint to retrieve server information
func (a *App) getServerIdentity() (*ServerIdentity, error) {
	endpoint := a.getAPIUrl("/identity")
	req, err := a.getRequest("GET", endpoint, nil)

	if err != nil {
		return nil, err
	}

	res, err := a.config.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		return nil, util.ErrorFromResponse(res)
	}

	var identity ServerIdentity

	if err = util.ReadResponseJSON(res, &identity); err != nil {
		return nil, err
	}

	return &identity, nil
}

// getServerCapabilities calls the / endpoint to retrieve server information
func (a *App) getServerCapabilities() (*ServerCapabilities, error) {
	endpoint := a.getAPIUrl("/")
	req, err := a.getRequest("GET", endpoint, nil)

	if err != nil {
		return nil, err
	}

	res, err := a.config.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		return nil, util.ErrorFromResponse(res)
	}

	var caps ServerCapabilities

	if err = util.ReadResponseJSON(res, &caps); err != nil {
		return nil, err
	}

	return &caps, nil
}

// getMe calls the /v2/user endpoint to retrieve the user the Plex token belongs to
func (a *App) getMe() (users.User, error) {
	var u users.User

	identity, err := a.getServerIdentity()

	if err != nil {
		return u, err
	}

	endpoint := a.getPlexAPIUrl("/user")
	req, err := a.getRequest("GET", endpoint, nil)

	if err != nil {
		return u, err
	}

	res, err := a.config.Client.Do(req)

	if err != nil {
		return u, err
	}

	if res.StatusCode/100 != 2 {
		return u, util.ErrorFromResponse(res)
	}

	var response User

	if err = util.ReadResponseJSON(res, &response); err != nil {
		return u, err
	}

	return plexUserToUser(response, identity.MediaContainer.MachineIdentifier), nil
}

func (a *App) getUsers() ([]users.User, error) {
	endpoint := a.getAPIUrl("/users")
	req, err := a.getRequest("GET", endpoint, nil)

	if err != nil {
		return nil, err
	}

	res, err := a.config.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		return nil, util.ErrorFromResponse(res)
	}

	var externalUsers ExternalUsers

	if err = util.ReadResponseJSON(res, &externalUsers); err != nil {
		return nil, err
	}

	return externalUserToUser(externalUsers), nil

}

// getPlexAPIUrl builds a URL for a Plex https://plex.tv/api/v2 endpoint
func (a *App) getPlexAPIUrl(elem ...string) string {
	path, err := url.JoinPath("https://plex.tv/api/v2", elem...)

	if err != nil {
		// Fatal since URLs need to be correct for the application to work as intended
		log.Fatalln(fmt.Sprintf("could not build API url for %q and %q: %e", a.config.URL, elem, err))
	}

	return path
}

// getAPIUrl builds a URL for a Plex endpoint using the
// configured base url
func (a *App) getAPIUrl(elem ...string) string {
	path, err := url.JoinPath(a.config.URL, elem...)

	if err != nil {
		// Fatal since URLs need to be correct for the application to work as intended
		log.Fatalln(fmt.Sprintf("could not build API url for %q and %q: %e", a.config.URL, elem, err))
	}

	return path
}

// getRequest is a helper method which adds the necessary headers
// to the request made to Plex
func (a *App) getRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, endpoint, body)

	if err != nil {
		return nil, err
	}

	r.Header.Set("X-Plex-Token", a.config.Token)
	r.Header.Set("Accept", "application/json,application/xml")

	return r, nil
}
