package app

const (
	MODE_DEBUG      = "debug"
	MODE_PRODUCTION = "production"
)

type oauthConfig struct {
	RedirectURL      string   `yaml:"redirect_url"`
	ClientID         string   `yaml:"client_id"`
	ClientSecret     string   `yaml:"client_secret"`
	Scopes           []string `yaml:"scopes"`
	UserInfoEndpoint string   `yaml:"user_info_endpoint"`
}

type conf struct {
	Mode             string `yaml:"mode"`
	Port             string `yaml:"port"`
	SessionSecret    string `yaml:"session_secret"`
	DSN              string `yaml:"dsn"`
	TimezoneLoc      string `yaml:"timezone_loc"`
	Summarize_ApiKey string `yaml:"Summarize_ApiKey"`
	OCR_ApiKey       string `yaml:"OCR_ApiKey"`
	ImagePath        string `yaml:"image_path"`
	OAuth2           struct {
		Google oauthConfig `yaml:"google"`
	} `yaml:"oauth2"`
}
