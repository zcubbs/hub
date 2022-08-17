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

		log.Println(groupCaption, tagCaption)
		fmt.Println(groups)
		//for _, group := range *groups {
		//	if group.Caption == groupCaption {
		//		for _, tag := range *group.Tags {
		//			if tag.Caption == tagCaption {
		//				return c.Render("tag", fiber.Map{
		//					"Title":        title,
		//					"SubTitle":     subTitle,
		//					"Logo":         logo,
		//					"Tag":          tag,
		//					"GroupCaption": groupCaption,
		//				})
		//			}
		//		}
		//	}
		//}

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
