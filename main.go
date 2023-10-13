package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/zcubbs/hub/internal/config"
	"html/template"
	"log"
	"runtime"
	"time"
)

var (
	Version = "0.0.0"
	Commit  = "none"
	Date    = "unknown"
)

var (
	configPath = flag.String("config", "./config.yaml", "Path to config file")
)

////go:embed web/public
//var publicFs embed.FS
//
////go:embed web/views
//var viewsFs embed.FS

func main() {
	flag.Parse()
	//Print out version information
	fmt.Println("Hub")
	fmt.Println("Version:", Version)
	fmt.Println("Commit:", Commit)
	fmt.Println("Date:", Date)
	fmt.Println(runtime.GOOS, runtime.GOARCH)

	// Load config
	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	groups := cfg.Data.Groups
	links := cfg.Data.Links
	footer := cfg.Data.Footer
	title := cfg.App.Title
	disclaimer := cfg.App.Disclaimer
	customHTML := cfg.App.CustomHtml
	showGithub := cfg.App.ShowGithub
	subTitle := cfg.App.Subtitle
	logo := cfg.App.LogoUrl
	devMode := cfg.Dev.Mode

	engine := html.New("./web/views", ".html")

	app := fiber.New(fiber.Config{
		Views:                 engine,
		ViewsLayout:           "layouts/main",
		DisableStartupMessage: true,
	})
	app.Static("/css", "./web/public/css", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})
	app.Static("/assets", "./web/public/assets", fiber.Static{
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
			"CustomHtml": template.HTML(customHTML), // unescaped html
			"Groups":     groups,
			"Links":      links,
			"Footer":     footer,
			"Version":    Version,
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
}
