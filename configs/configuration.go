package configs

type Configuration struct {
	Server `mapstructure:"server" json:"server"`
	App    `mapstructure:"app" json:"app"`
	Dev    `mapstructure:"dev" json:"dev"`
	Groups *[]Group `mapstructure:"groups" json:"groups"`
	Links  *[]Link  `mapstructure:"links" json:"links"`
}

type Group struct {
	Id      string  `mapstructure:"id" json:"id"`
	Caption string  `mapstructure:"caption" json:"caption"`
	Links   *[]Link `mapstructure:"links" json:"links"`
	Hidden  bool    `mapstructure:"hidden" json:"hidden"`
}

type Link struct {
	Caption  string `mapstructure:"caption" json:"caption"`
	URL      string `mapstructure:"url" json:"url"`
	Icon     string `mapstructure:"icon" json:"icon"`
	NewTab   bool   `mapstructure:"newTab" json:"newTab"`
	External bool   `mapstructure:"external" json:"external"`
}

type App struct {
	CustomHtml string `mapstructure:"customHtml" json:"customHtml"`
	Title      string `mapstructure:"title" json:"title"`
	Subtitle   string `mapstructure:"subtitle" json:"subtitle"`
	ShowGithub bool   `mapstructure:"showGithub" json:"showGithub"`
	LogoUrl    string `mapstructure:"logoUrl" json:"logoUrl"`
	Disclaimer string `mapstructure:"disclaimer" json:"disclaimer"`
}

type Server struct {
	Port   string `mapstructure:"port" json:"port"`
	Secure bool   `mapstructure:"secure" json:"secure"`
	Tz     string `mapstructure:"tz" json:"tz"`
}

type Dev struct {
	Mode bool `mapstructure:"mode" json:"mode"`
}
