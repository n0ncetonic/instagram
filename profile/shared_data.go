package profile

// SharedData contains all data shared by a profile
type SharedData struct {
	ActivityCounts *string     `json:"activity_counts"`
	Config         Config      `json:"config"`
	SupportsES6    bool        `json:"supports_es6"`
	CountryCode    string      `json:"country_code"`
	LanguageCode   string      `json:"language_code"`
	Locale         string      `json:"locale"`
	EntryData      EntryData   `json:"entry_data"`
	Gatekeepers    Gatekeepers `json:"gatekeepers"`
	Knobs          Knobs       `json:"knobs"`
	Qe             Qe          `json:"qe"`
	Hostname       string      `json:"hostname"`
	Platform       string      `json:"platform"`
	RhxGis         string      `json:"rhx_gis"`
	Nonce          string      `json:"nonce"`
	ZeroData       ZeroData    `json:"zero_data"`
	RolloutHash    string      `json:"rollout_hash"`
	BundleVariant  string      `json:"bundle_variant"`
	ProbablyHasApp bool        `json:"probably_has_app"`
	ShowAppInstall bool        `json:"show_app_install"`
}

// Config holds a token and viewer information used by Instagram
type Config struct {
	CsrfToken string  `json:"csrf_token"`
	Viewer    *string `json:"viewer"`
}

// EntryData wraps a ProfilePage array
type EntryData struct {
	ProfilePage []ProfilePage `json:"ProfilePage"`
}

// ProfilePage wraps the GraphQL for a profile
type ProfilePage struct {
	LoggingPageID                 string                   `json:"logging_page_id"`
	ShowSuggestedProfiles         bool                     `json:"show_suggested_profiles"`
	GraphQL                       GraphQL                  `json:"graphql"`
	FelixOnboardingVideoResources OnboardingVideoResources `json:"felix_onboarding_video_resources"`
}

// GraphQL wraps our main User struct
type GraphQL struct {
	User User `json:"user"`
}

// OnboardingVideoResources holds onboarding(?) video information
type OnboardingVideoResources struct {
	Mp4    string `json:"mp4"`
	Poster string `json:"poster"`
}

// Gatekeepers holds what appear to be SEO related options
type Gatekeepers struct {
	Ld    bool `json:"ld"`
	Seo   bool `json:"seo"`
	Seoht bool `json:"seoht"`
	Sf    bool `json:"sf"`
	Saa   bool `json:"saa"`
}

// Knobs contains what appear to be security controls
type Knobs struct {
	Acctntb int `json:"acct:ntb"`
	Cb      int `json:"cb"`
	Captcha int `json:"captcha"`
}

// Qe holds an arbitrary amount of what appear to be page rendering options
type Qe struct {
	Qe []interface{} `json:"qe"`
}

// ZeroData is an empty struct
type ZeroData struct {
}
