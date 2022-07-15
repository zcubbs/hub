package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/zcubbs/hub/internal/models"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	groups := loadGroupsFromYaml()
	title := getTitle()
	disclaimer := getDisclaimer()
	customHTML := getCustomHtml()
	showGithub := getShowGithub()
	subTitle := getSubTitle()
	logo := getLogo()
	engine := html.New("./internal/views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Static("/css", "./internal/public/css", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})
	app.Static("/assets", "./internal/public/assets", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	engine.Reload(isDevMode())

	log.Printf("Custom HTML: %s\n", customHTML)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title":      title,
			"SubTitle":   subTitle,
			"Logo":       logo,
			"Disclaimer": disclaimer,
			"ShowGithub": showGithub,
			"CustomHtml": template.HTML(customHTML),
			"Groups":     groups,
		})
	})

	app.Get("/tag/:group/:caption", func(c *fiber.Ctx) error {
		groupCaption, err := url.QueryUnescape(c.Params("group"))
		if err != nil {
			log.Println(err)
		}
		tagCaption, err := url.QueryUnescape(c.Params("caption"))
		if err != nil {
			log.Println(err)
		}
		for _, group := range *groups {
			if group.Caption == groupCaption {
				for _, tag := range *group.Tags {
					if tag.Caption == tagCaption {
						return c.Render("tag", fiber.Map{
							"Title":        title,
							"SubTitle":     subTitle,
							"Logo":         logo,
							"Tag":          tag,
							"GroupCaption": groupCaption,
						})
					}
				}
			}
		}

		return c.Render("index", fiber.Map{
			"Title":      title,
			"SubTitle":   subTitle,
			"Logo":       logo,
			"Disclaimer": disclaimer,
			"ShowGithub": showGithub,
			"CustomHtml": template.HTML(customHTML),
			"Tags":       groups,
		})

	})

	log.Fatal(app.Listen(":8000"))
}

func loadGroupsFromYaml() *[]models.Group {
	yamlFile, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	yamlGroups := &models.Groups{}

	err = yaml.Unmarshal(yamlFile, yamlGroups)
	if err != nil {
		log.Fatal(err)
	}

	return yamlGroups.Groups
}

func getConfigPath() string {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}
	return configPath
}

func isDevMode() bool {
	devMode := os.Getenv("DEV_MODE")
	if devMode == "" {
		devMode = "false"
	}
	boolVal, _ := strconv.ParseBool(devMode)
	return boolVal
}

func getDisclaimer() string {
	appTitle := os.Getenv("APP_DISCLAIMER")
	if appTitle == "" {
		appTitle = "Link Hub For everyone."
	}
	return appTitle
}

func getCustomHtml() string {
	customHtml := os.Getenv("APP_CUSTOM_HTML")
	if customHtml == "" {
		customHtml = "<span></span>"
	}
	return customHtml
}

func getTitle() string {
	appTitle := os.Getenv("APP_TITLE")
	if appTitle == "" {
		appTitle = "z/HuB"
	}
	return appTitle
}

func getShowGithub() bool {
	showGithub := os.Getenv("SHOW_GITHUB")
	if showGithub == "false" {
		return false
	}
	return true
}

func getLogo() string {
	logo := os.Getenv("LOGO_URL")
	if logo == "" {
		return "/assets/logo.png"
	}
	return logo
}

func getSubTitle() string {
	subTitle := os.Getenv("APP_SUB_TITLE")
	if subTitle == "" {
		return "HuB"
	}
	return subTitle
}
