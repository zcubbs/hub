package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	"github.com/zcubbs/hub/configs"
	"html/template"
	"log"
	"net/url"
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

	groups := configs.Config.Groups
	links := configs.Config.Links
	title := configs.Config.App.Title
	disclaimer := configs.Config.App.Disclaimer
	customHTML := configs.Config.App.CustomHtml
	showGithub := configs.Config.App.ShowGithub
	subTitle := configs.Config.App.Subtitle
	logo := configs.Config.App.LogoUrl
	devMode := configs.Config.Dev.Mode
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
			"IsSubGroup": false,
		})
	})

	app.Get("/group/:groupId/:subGroupId", func(c *fiber.Ctx) error {
		groupId, err := url.QueryUnescape(c.Params("groupId"))
		if err != nil {
			log.Println(err)
		}
		subGroupId, err := url.QueryUnescape(c.Params("subGroupId"))
		if err != nil {
			log.Println(err)
		}

		groupCaption := ""
		subGroupCaption := ""

		for _, group := range *groups {
			if group.Id == groupId {
				groupCaption = group.Caption
			}
			if group.Id == subGroupId {
				subGroupCaption = group.Caption
			}
		}

		for _, group := range *groups {
			if group.Id == subGroupId {
				return c.Render("group", fiber.Map{
					"Title":           title,
					"SubTitle":        subTitle,
					"Logo":            logo,
					"Links":           group.Links,
					"GroupCaption":    groupCaption,
					"SubGroupCaption": subGroupCaption,
					"IsSubGroup":      true,
				})
			}
		}

		return c.Render("group", fiber.Map{
			"Title":      title,
			"SubTitle":   subTitle,
			"Logo":       logo,
			"Disclaimer": disclaimer,
			"ShowGithub": showGithub,
			"CustomHtml": template.HTML(customHTML),
			"Groups":     groups,
			"Links":      links,
			"IsSubGroup": false,
		})

	})

	log.Fatal(app.Listen(":8000"))
}
