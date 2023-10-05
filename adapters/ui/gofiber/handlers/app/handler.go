package app

import (
	"os"

	_ "github.com/rof20004/healthy-api/adapters/ui/gofiber/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func InitAppRoutes(app *fiber.App) {
	app.Get("/health", health)
	app.Get("/swagger", swagger.HandlerDefault)

	var (
		url               = "https://api.ajuda.academy/swagger/doc.json"
		oauth2RedirectUrl = "https://api.ajuda.academy/swagger/oauth2-redirect.html"
	)

	if os.Getenv("ENV") == "local" {
		url = "http://localhost:8080/swagger/doc.json"
		oauth2RedirectUrl = "http://localhost:8080/swagger/oauth2-redirect.html"
	}

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         url,
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: oauth2RedirectUrl,
	}))
}

func health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"health": "ok",
	})
}
