package jellyfin

import (
	"errors"
	"fmt"
	"gobble/pkg"
	"gobble/pkg/common"
	"gobble/pkg/users"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// App represents a Jellyfin media app
type App struct {
	config Config
}

// Config represents a config file entry for a Jellyfin app
type Config struct {
	URL    string `json:"url" toml:"url"`
	Token  string `json:"token" toml:"token"`
	Client *http.Client
}

// New creates the provided config into a new Jellyfin app
func New(config *Config) *App {
	if config.Client == nil {
		config.Client = &http.Client{Timeout: 30 * time.Second}
	}

	return &App{
		config: *config,
	}
}

// GetUsers retrieves all users from the Jellyfin app,
// mapped to a normalized User object
func (j *App) GetUsers() ([]users.User, error) {
	endpoint := j.getAPIUrl("/Users")
	req, err := j.getRequest("GET", endpoint, nil)

	if err != nil {
		return nil, err
	}

	res, err := j.config.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode/100 != 2 {
		return nil, errors.New(fmt.Sprintf("request to %q unsuccessful: %d %s", endpoint, res.StatusCode, res.Status))
	}

	var content []User

	if err = pkg.ReadResponseJSON(res, &content); err != nil {
		return nil, err
	}

	return convertUsers(content), nil
}

// GetSystemInfo retrieves system info from the Jellyfin API
func (j *App) GetSystemInfo() (*common.SystemInfo, error) {
	endpoint := j.getAPIUrl("/System/Info")
	req, err := j.getRequest("GET", endpoint, nil)

	if err != nil {
		return nil, err
	}

	res, err := j.config.Client.Do(req)

	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("request to %q unsuccessful: %d %s", endpoint, res.StatusCode, res.Status)
	}
	var info SystemInfo

	if err = pkg.ReadResponseJSON(res, &info); err != nil {
		return nil, err
	}

	return convertSystemInfo(info), nil
}

// getAPIUrl builds a URL for a Jellyfin endpoint using the
// configured base url
func (j *App) getAPIUrl(elem ...string) string {
	path, err := url.JoinPath(j.config.URL, elem...)

	if err != nil {
		// Fatal since URLs need to be correct for the application to work as intended
		log.Fatalln(fmt.Sprintf("could not build API url for %q and %q: %e", j.config.URL, elem, err))
	}

	return path
}

// getRequest is a helper method which adds the necessary headers
// to the request made to Jellyfin
func (j *App) getRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, endpoint, body)

	if err != nil {
		return nil, err
	}

	r.Header.Set("X-Emby-Token", j.config.Token)
	r.Header.Set("Accept", "application/json")

	return r, nil
}
