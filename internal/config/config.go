package config

type Configuration struct {
	Server `mapstructure:"server" json:"server"`
	App    `mapstructure:"app" json:"app"`
	Dev    `mapstructure:"dev" json:"dev"`
	Data   `mapstructure:"data" json:"data"`
}

type Group struct {
	Caption  string    `mapstructure:"caption" json:"caption"`
	Links    []Link    `mapstructure:"links" json:"links"`
	Sections []Section `mapstructure:"sections" json:"sections"`
}

type Data struct {
	Groups []Group `mapstructure:"groups" json:"groups"`
	Links  []Link  `mapstructure:"links" json:"links"`
	Footer Footer  `mapstructure:"footerLinks" json:"footerLinks"`
}

type Section struct {
	Caption string `mapstructure:"caption" json:"caption"`
	Links   []Link `mapstructure:"links" json:"links"`
}

type Link struct {
	Caption string `mapstructure:"caption" json:"caption"`
	URL     string `mapstructure:"url" json:"url"`
	Icon    string `mapstructure:"icon" json:"icon"`
	NewTab  bool   `mapstructure:"newTab" json:"newTab"`
	Links   []Link `mapstructure:"links" json:"links"`
}

type App struct {
	CustomHtml string `mapstructure:"customHtml" json:"customHtml"`
	Title      string `mapstructure:"title" json:"title"`
	Subtitle   string `mapstructure:"subtitle" json:"subtitle"`
	ShowGithub bool   `mapstructure:"showGithub" json:"showGithub"`
	LogoUrl    string `mapstructure:"logoUrl" json:"logoUrl"`
	Disclaimer string `mapstructure:"disclaimer" json:"disclaimer"`
	Debug      bool   `mapstructure:"debug" json:"debug"`
}

type Server struct {
	Port   int    `mapstructure:"port" json:"port"`
	Secure bool   `mapstructure:"secure" json:"secure"`
	Tz     string `mapstructure:"tz" json:"tz"`
}

type Dev struct {
	Mode bool `mapstructure:"mode" json:"mode"`
}

type Footer struct {
	Links []Link `mapstructure:"links" json:"links"`
}
