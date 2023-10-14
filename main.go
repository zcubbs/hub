package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
	"github.com/zcubbs/hub/internal/config"
	"html/template"
	"log"
	"net/http"
	"runtime"
)

var (
	Version = "0.0.0"
	Commit  = "none"
	Date    = "unknown"
)

var (
	configPath = flag.String("config", "./config.yaml", "Path to config file")
)

//go:embed web/public/css/*
var cssFs embed.FS

//go:embed web/public/assets/*
var assetsFs embed.FS

//go:embed web/views/*
var viewsFs embed.FS

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

	// init template engine
	engine := html.NewFileSystem(http.FS(viewsFs), ".html")
	engine.Engine.Directory = "web/views"
	engine.Reload(devMode)

	// init server
	app := fiber.New(fiber.Config{
		Views:                 engine,
		ViewsLayout:           "layouts/main",
		DisableStartupMessage: true,
	})

	// serve static files
	app.Use("/css", filesystem.New(
		filesystem.Config{
			Root:       http.FS(cssFs),
			PathPrefix: "web/public/css",
			Browse:     false,
		},
	))
	app.Use("/assets", filesystem.New(
		filesystem.Config{
			Root:       http.FS(assetsFs),
			PathPrefix: "web/public/assets",
			Browse:     false,
		},
	))

	// serve index
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

	// serve
	log.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
}
