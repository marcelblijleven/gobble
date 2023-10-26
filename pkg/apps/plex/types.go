package plex

import "encoding/xml"

// ExternalUsers represents the Plex API response for the /users endpoint
type ExternalUsers struct {
	XMLName           xml.Name       `xml:"MediaContainer"`
	Text              string         `xml:",chardata"`
	FriendlyName      string         `xml:"friendlyName,attr"`
	Identifier        string         `xml:"identifier,attr"`
	MachineIdentifier string         `xml:"machineIdentifier,attr"`
	TotalSize         string         `xml:"totalSize,attr"`
	Size              string         `xml:"size,attr"`
	User              []externalUser `xml:"User"`
}

// externalUser is part of the ExternalUsers response
type externalUser struct {
	Text                      string `xml:",chardata"`
	ID                        string `xml:"id,attr"`
	Title                     string `xml:"title,attr"`
	Username                  string `xml:"username,attr"`
	Email                     string `xml:"email,attr"`
	RecommendationsPlaylistId string `xml:"recommendationsPlaylistId,attr"`
	Thumb                     string `xml:"thumb,attr"`
	Protected                 string `xml:"protected,attr"`
	Home                      string `xml:"home,attr"`
	AllowTuners               string `xml:"allowTuners,attr"`
	AllowSync                 string `xml:"allowSync,attr"`
	AllowCameraUpload         string `xml:"allowCameraUpload,attr"`
	AllowChannels             string `xml:"allowChannels,attr"`
	AllowSubtitleAdmin        string `xml:"allowSubtitleAdmin,attr"`
	FilterAll                 string `xml:"filterAll,attr"`
	FilterMovies              string `xml:"filterMovies,attr"`
	FilterMusic               string `xml:"filterMusic,attr"`
	FilterPhotos              string `xml:"filterPhotos,attr"`
	FilterTelevision          string `xml:"filterTelevision,attr"`
	Restricted                string `xml:"restricted,attr"`
}

// User represents the Plex API response for the /v2/user endpoint
type User struct {
	Id                      int           `json:"id"`
	Uuid                    string        `json:"uuid"`
	Username                string        `json:"username"`
	Title                   string        `json:"title"`
	Email                   string        `json:"email"`
	FriendlyName            string        `json:"friendlyName"`
	Locale                  interface{}   `json:"locale"`
	Confirmed               bool          `json:"confirmed"`
	JoinedAt                int           `json:"joinedAt"`
	EmailOnlyAuth           bool          `json:"emailOnlyAuth"`
	HasPassword             bool          `json:"hasPassword"`
	Protected               bool          `json:"protected"`
	Thumb                   string        `json:"thumb"`
	AuthToken               string        `json:"authToken"`
	MailingListStatus       string        `json:"mailingListStatus"`
	MailingListActive       bool          `json:"mailingListActive"`
	ScrobbleTypes           string        `json:"scrobbleTypes"`
	Country                 string        `json:"country"`
	Subscription            subscription  `json:"subscription"`
	SubscriptionDescription interface{}   `json:"subscriptionDescription"`
	Restricted              bool          `json:"restricted"`
	Anonymous               bool          `json:"anonymous"`
	Home                    bool          `json:"home"`
	Guest                   bool          `json:"guest"`
	HomeSize                int           `json:"homeSize"`
	HomeAdmin               bool          `json:"homeAdmin"`
	MaxHomeSize             int           `json:"maxHomeSize"`
	RememberExpiresAt       int           `json:"rememberExpiresAt"`
	Profile                 profile       `json:"profile"`
	Entitlements            []interface{} `json:"entitlements"`
	Services                []service     `json:"apps"`
	AdsConsent              interface{}   `json:"adsConsent"`
	AdsConsentSetAt         interface{}   `json:"adsConsentSetAt"`
	AdsConsentReminderAt    interface{}   `json:"adsConsentReminderAt"`
	ExperimentalFeatures    bool          `json:"experimentalFeatures"`
	TwoFactorEnabled        bool          `json:"twoFactorEnabled"`
	BackupCodesCreated      bool          `json:"backupCodesCreated"`
}

// subscription is part of the User struct
type subscription struct {
	Active         bool        `json:"active"`
	SubscribedAt   interface{} `json:"subscribedAt"`
	Status         string      `json:"status"`
	PaymentService interface{} `json:"paymentService"`
	Plan           interface{} `json:"plan"`
	Features       []string    `json:"features"`
}

// profile is part of the User struct
type profile struct {
	AutoSelectAudio              bool   `json:"autoSelectAudio"`
	DefaultAudioLanguage         string `json:"defaultAudioLanguage"`
	DefaultSubtitleLanguage      string `json:"defaultSubtitleLanguage"`
	AutoSelectSubtitle           int    `json:"autoSelectSubtitle"`
	DefaultSubtitleAccessibility int    `json:"defaultSubtitleAccessibility"`
	DefaultSubtitleForced        int    `json:"defaultSubtitleForced"`
}

// service is part of the User struct
type service struct {
	Identifier string  `json:"identifier"`
	Endpoint   string  `json:"endpoint"`
	Token      *string `json:"token"`
	Secret     *string `json:"secret"`
	Status     string  `json:"status"`
}

// ServerIdentity represents the response from the Plex API /identity endpoint
type ServerIdentity struct {
	MediaContainer struct {
		Size              int    `json:"size"`
		Claimed           bool   `json:"claimed"`
		MachineIdentifier string `json:"machineIdentifier"`
		Version           string `json:"version"`
	} `json:"MediaContainer"`
}

// ServerCapabilities represents the response from the Plex API / endpoint
type ServerCapabilities struct {
	MediaContainer struct {
		Size                          int         `json:"size"`
		AllowCameraUpload             bool        `json:"allowCameraUpload"`
		AllowChannelAccess            bool        `json:"allowChannelAccess"`
		AllowMediaDeletion            bool        `json:"allowMediaDeletion"`
		AllowSharing                  bool        `json:"allowSharing"`
		AllowSync                     bool        `json:"allowSync"`
		AllowTuners                   bool        `json:"allowTuners"`
		BackgroundProcessing          bool        `json:"backgroundProcessing"`
		Certificate                   bool        `json:"certificate"`
		CompanionProxy                bool        `json:"companionProxy"`
		CountryCode                   string      `json:"countryCode"`
		Diagnostics                   string      `json:"diagnostics"`
		EventStream                   bool        `json:"eventStream"`
		FriendlyName                  string      `json:"friendlyName"`
		HubSearch                     bool        `json:"hubSearch"`
		ItemClusters                  bool        `json:"itemClusters"`
		Livetv                        int         `json:"livetv"`
		MachineIdentifier             string      `json:"machineIdentifier"`
		MediaProviders                bool        `json:"mediaProviders"`
		Multiuser                     bool        `json:"multiuser"`
		MusicAnalysis                 int         `json:"musicAnalysis"`
		MyPlex                        bool        `json:"myPlex"`
		MyPlexMappingState            string      `json:"myPlexMappingState"`
		MyPlexSigninState             string      `json:"myPlexSigninState"`
		MyPlexSubscription            bool        `json:"myPlexSubscription"`
		MyPlexUsername                string      `json:"myPlexUsername"`
		OfflineTranscode              int         `json:"offlineTranscode"`
		OwnerFeatures                 string      `json:"ownerFeatures"`
		PhotoAutoTag                  bool        `json:"photoAutoTag"`
		Platform                      string      `json:"platform"`
		PlatformVersion               string      `json:"platformVersion"`
		PluginHost                    bool        `json:"pluginHost"`
		PushNotifications             bool        `json:"pushNotifications"`
		ReadOnlyLibraries             bool        `json:"readOnlyLibraries"`
		StreamingBrainABRVersion      int         `json:"streamingBrainABRVersion"`
		StreamingBrainVersion         int         `json:"streamingBrainVersion"`
		Sync                          bool        `json:"sync"`
		TranscoderActiveVideoSessions int         `json:"transcoderActiveVideoSessions"`
		TranscoderAudio               bool        `json:"transcoderAudio"`
		TranscoderLyrics              bool        `json:"transcoderLyrics"`
		TranscoderPhoto               bool        `json:"transcoderPhoto"`
		TranscoderSubtitles           bool        `json:"transcoderSubtitles"`
		TranscoderVideo               bool        `json:"transcoderVideo"`
		TranscoderVideoBitrates       string      `json:"transcoderVideoBitrates"`
		TranscoderVideoQualities      string      `json:"transcoderVideoQualities"`
		TranscoderVideoResolutions    string      `json:"transcoderVideoResolutions"`
		UpdatedAt                     int         `json:"updatedAt"`
		Updater                       bool        `json:"updater"`
		Version                       string      `json:"version"`
		VoiceSearch                   bool        `json:"voiceSearch"`
		Directory                     []directory `json:"Directory"`
	} `json:"MediaContainer"`
}

// directory is part of the ServerCapabilities endpoint
type directory struct {
	Count int    `json:"count"`
	Key   string `json:"key"`
	Title string `json:"title"`
}
