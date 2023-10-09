package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/zcubbs/hub/configs"
	"html/template"
	"log"
	"time"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	//Print out splash screen
	fmt.Printf(`%s%-16s`+"\n\n", configs.Splash, configs.Version)

	//Bootstrap configs
	configs.Bootstrap()

	groups := configs.Config.Data.Groups
	links := configs.Config.Data.Links
	title := configs.Config.App.Title
	disclaimer := configs.Config.App.Disclaimer
	customHTML := configs.Config.App.CustomHtml
	showGithub := configs.Config.App.ShowGithub
	subTitle := configs.Config.App.Subtitle
	logo := configs.Config.App.LogoUrl
	devMode := configs.Config.Dev.Mode
	debugMode := configs.Config.App.Debug

	if debugMode {
		configs.DebugConfig()
	}

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

	engine.Reload(devMode)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("group", fiber.Map{
			"Title":      title,
			"SubTitle":   subTitle,
			"Logo":       logo,
			"Disclaimer": disclaimer,
			"ShowGithub": showGithub,
			"CustomHtml": template.HTML(customHTML),
			"Groups":     groups,
			"Links":      links,
		})
	})

	log.Fatal(app.Listen(":8000"))
}
