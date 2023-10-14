package config

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

type AppConfiguration interface {
	*Configuration
}

func Load(path string) (*Configuration, error) {
	var cfg Configuration

	var onceEnv sync.Once
	onceEnv.Do(loadEnv) // load env only once

	initViperPresets(path)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("warn: unable to load config file path=%s err=%s\n", path, err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("could not decode config into struct err=%s", err)
	}

	if err := validate(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed err=%w", err)
	}

	if cfg.App.Debug {
		print(&cfg)
	}

	return &cfg, nil
}

func initViperPresets(path string) {
	dir := filepath.Dir(path)
	file := filepath.Base(path)
	viper.AddConfigPath(dir)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
}

func validate(_ *Configuration) error {
	return nil
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("no .env file found")
	}
}

func print(recipe *Configuration) {
	jsonConfig, err := json.MarshalIndent(recipe, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("config path: %s\n", viper.ConfigFileUsed())
	fmt.Printf("%v\n", string(jsonConfig))
}

var defaults = map[string]interface{}{
	"server.secure":  false,
	"server.port":    8000,
	"server.tz":      "Europe/Paris",
	"app.customHtml": `<span></span>`,
	"app.disclaimer": `Link Hub For everyone.`,
	"app.title":      "Hub",
	"app.subtitle":   "Developer Hub",
	"app.showGithub": false,
	"app.logoUrl":    "/assets/logo.png",
	"data.links":     []Link{},
	"data.groups":    []Group{},
	"data.footer":    Footer{},
}
