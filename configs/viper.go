package configs

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"

	"github.com/spf13/viper"
	"log"
	"strings"
)

var cfgFile string

var Config Configuration

type AppConfiguration interface {
	*Configuration
}

var (
	defaults = map[string]interface{}{
		"server.secure":  false,
		"server.port":    "8000",
		"server.tz":      "Europe/Paris",
		"app.customHtml": `<span></span>`,
		"app.disclaimer": `Link Hub For everyone.`,
		"app.title":      "Hub",
		"app.subtitle":   "Developer Hub",
		"app.showGithub": true,
		"app.logoUrl":    "/assets/logo.png",
	}
	allowedEnvVarKeys = []string{
		"app.customHtml",
		"app.title",
		"app.subtitle",
		"app.showGithub",
		"app.logoUrl",
		"app.disclaimer",
		"dev.mode",
		"dev.debug",
		"server.tz",
		"server.secure",
		"server.port",
	}
	envPrefix   = "HUB"
	configName  = "config"
	configType  = "yaml"
	configPaths = []string{
		".",
	}
)

// Bootstrap reads in config file and ENV variables if set.
func Bootstrap() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		for _, p := range configPaths {
			viper.AddConfigPath(p)
		}
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}
	}
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(envPrefix)

	for _, key := range allowedEnvVarKeys {
		err := viper.BindEnv(key)
		if err != nil {
			fmt.Println(err)
		}
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("could not decode config into struct: %v", err)
	}

	log.Println("Configuration loaded")
	//debugConfig()
}

func debugConfig() {
	jsonConfig, err := json.MarshalIndent(&Config, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", string(jsonConfig))
}
