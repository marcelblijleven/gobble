package jellyfin

import "time"

// User is the Jellyfin response type for a user
type User struct {
	Name                      string        `json:"Name"`
	ServerId                  string        `json:"ServerId"`
	ServerName                string        `json:"ServerName"`
	ID                        string        `json:"Id"`
	PrimaryImageTag           string        `json:"PrimaryImageTag"`
	HasPassword               bool          `json:"HasPassword"`
	HasConfiguredPassword     bool          `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword bool          `json:"HasConfiguredEasyPassword"`
	EnableAutoLogin           bool          `json:"EnableAutoLogin"`
	LastLoginDate             time.Time     `json:"LastLoginDate"`
	LastActivityDate          time.Time     `json:"LastActivityDate"`
	Configuration             configuration `json:"Configuration"`
	Policy                    policy        `json:"Policy"`
	PrimaryImageAspectRatio   int           `json:"PrimaryImageAspectRatio"`
}

// configuration is part of the User response type
type configuration struct {
	AudioLanguagePreference    string   `json:"AudioLanguagePreference"`
	PlayDefaultAudioTrack      bool     `json:"PlayDefaultAudioTrack"`
	SubtitleLanguagePreference string   `json:"SubtitleLanguagePreference"`
	DisplayMissingEpisodes     bool     `json:"DisplayMissingEpisodes"`
	GroupedFolders             []string `json:"GroupedFolders"`
	SubtitleMode               string   `json:"SubtitleMode"`
	DisplayCollectionsView     bool     `json:"DisplayCollectionsView"`
	EnableLocalPassword        bool     `json:"EnableLocalPassword"`
	OrderedViews               []string `json:"OrderedViews"`
	LatestItemsExcludes        []string `json:"LatestItemsExcludes"`
	MyMediaExcludes            []string `json:"MyMediaExcludes"`
	HidePlayedInLatest         bool     `json:"HidePlayedInLatest"`
	RememberAudioSelections    bool     `json:"RememberAudioSelections"`
	RememberSubtitleSelections bool     `json:"RememberSubtitleSelections"`
	EnableNextEpisodeAutoPlay  bool     `json:"EnableNextEpisodeAutoPlay"`
}

// policy is part of the User response type
type policy struct {
	IsAdministrator                  bool       `json:"IsAdministrator"`
	IsHidden                         bool       `json:"IsHidden"`
	IsDisabled                       bool       `json:"IsDisabled"`
	MaxParentalRating                int        `json:"MaxParentalRating"`
	BlockedTags                      []string   `json:"BlockedTags"`
	EnableUserPreferenceAccess       bool       `json:"EnableUserPreferenceAccess"`
	AccessSchedules                  []schedule `json:"AccessSchedules"`
	BlockUnratedItems                []string   `json:"BlockUnratedItems"`
	EnableRemoteControlOfOtherUsers  bool       `json:"EnableRemoteControlOfOtherUsers"`
	EnableSharedDeviceControl        bool       `json:"EnableSharedDeviceControl"`
	EnableRemoteAccess               bool       `json:"EnableRemoteAccess"`
	EnableLiveTvManagement           bool       `json:"EnableLiveTvManagement"`
	EnableLiveTvAccess               bool       `json:"EnableLiveTvAccess"`
	EnableMediaPlayback              bool       `json:"EnableMediaPlayback"`
	EnableAudioPlaybackTranscoding   bool       `json:"EnableAudioPlaybackTranscoding"`
	EnableVideoPlaybackTranscoding   bool       `json:"EnableVideoPlaybackTranscoding"`
	EnablePlaybackRemuxing           bool       `json:"EnablePlaybackRemuxing"`
	ForceRemoteSourceTranscoding     bool       `json:"ForceRemoteSourceTranscoding"`
	EnableContentDeletion            bool       `json:"EnableContentDeletion"`
	EnableContentDeletionFromFolders []string   `json:"EnableContentDeletionFromFolders"`
	EnableContentDownloading         bool       `json:"EnableContentDownloading"`
	EnableSyncTranscoding            bool       `json:"EnableSyncTranscoding"`
	EnableMediaConversion            bool       `json:"EnableMediaConversion"`
	EnabledDevices                   []string   `json:"EnabledDevices"`
	EnableAllDevices                 bool       `json:"EnableAllDevices"`
	EnabledChannels                  []string   `json:"EnabledChannels"`
	EnableAllChannels                bool       `json:"EnableAllChannels"`
	EnabledFolders                   []string   `json:"EnabledFolders"`
	EnableAllFolders                 bool       `json:"EnableAllFolders"`
	InvalidLoginAttemptCount         int        `json:"InvalidLoginAttemptCount"`
	LoginAttemptsBeforeLockout       int        `json:"LoginAttemptsBeforeLockout"`
	MaxActiveSessions                int        `json:"MaxActiveSessions"`
	EnablePublicSharing              bool       `json:"EnablePublicSharing"`
	BlockedMediaFolders              []string   `json:"BlockedMediaFolders"`
	BlockedChannels                  []string   `json:"BlockedChannels"`
	RemoteClientBitrateLimit         int        `json:"RemoteClientBitrateLimit"`
	AuthenticationProviderId         string     `json:"AuthenticationProviderId"`
	PasswordResetProviderId          string     `json:"PasswordResetProviderId"`
	SyncPlayAccess                   string     `json:"SyncPlayAccess"`
}

// schedule is part of the policy response type
type schedule struct {
	Id        int    `json:"Id"`
	UserId    string `json:"UserId"`
	DayOfWeek string `json:"DayOfWeek"`
	StartHour int    `json:"StartHour"`
	EndHour   int    `json:"EndHour"`
}

// SystemInfo is the Jellyfin API response for /System/Info
type SystemInfo struct {
	LocalAddress               string                  `json:"LocalAddress"`
	ServerName                 string                  `json:"ServerName"`
	Version                    string                  `json:"Version"`
	ProductName                string                  `json:"ProductName"`
	OperatingSystem            string                  `json:"OperatingSystem"`
	Id                         string                  `json:"Id"`
	StartupWizardCompleted     bool                    `json:"StartupWizardCompleted"`
	OperatingSystemDisplayName string                  `json:"OperatingSystemDisplayName"`
	PackageName                string                  `json:"PackageName"`
	HasPendingRestart          bool                    `json:"HasPendingRestart"`
	IsShuttingDown             bool                    `json:"IsShuttingDown"`
	SupportsLibraryMonitor     bool                    `json:"SupportsLibraryMonitor"`
	WebSocketPortNumber        int                     `json:"WebSocketPortNumber"`
	CompletedInstallations     []completedInstallation `json:"CompletedInstallations"`
	CanSelfRestart             bool                    `json:"CanSelfRestart"`
	CanLaunchWebBrowser        bool                    `json:"CanLaunchWebBrowser"`
	ProgramDataPath            string                  `json:"ProgramDataPath"`
	WebPath                    string                  `json:"WebPath"`
	ItemsByNamePath            string                  `json:"ItemsByNamePath"`
	CachePath                  string                  `json:"CachePath"`
	LogPath                    string                  `json:"LogPath"`
	InternalMetadataPath       string                  `json:"InternalMetadataPath"`
	TranscodingTempPath        string                  `json:"TranscodingTempPath"`
	SystemArchitecture         string                  `json:"SystemArchitecture"`
}

// completedInstallation is part of the SystemInfo response
type completedInstallation struct {
	Guid        string      `json:"Guid"`
	Name        string      `json:"Name"`
	Version     string      `json:"Version"`
	Changelog   string      `json:"Changelog"`
	SourceUrl   string      `json:"SourceUrl"`
	Checksum    string      `json:"Checksum"`
	PackageInfo packageInfo `json:"PackageInfo"`
}

// packageInfo is part of the completedInstallation
type packageInfo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Overview    string    `json:"overview"`
	Owner       string    `json:"owner"`
	Category    string    `json:"category"`
	Guid        string    `json:"guid"`
	Versions    []version `json:"versions"`
	ImageUrl    string    `json:"imageUrl"`
}

// version is part of the packageInfo
type version struct {
	Version        string `json:"version"`
	VersionNumber  string `json:"VersionNumber"`
	Changelog      string `json:"changelog"`
	TargetAbi      string `json:"targetAbi"`
	SourceUrl      string `json:"sourceUrl"`
	Checksum       string `json:"checksum"`
	Timestamp      string `json:"timestamp"`
	RepositoryName string `json:"repositoryName"`
	RepositoryUrl  string `json:"repositoryUrl"`
}
