package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	appHandler "github.com/rof20004/healthy-api/ui/presentation/gofiber/handlers/app"
	userHandler "github.com/rof20004/healthy-api/ui/presentation/gofiber/handlers/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	appHandler.InitAppRoutes(app)
	userHandler.InitUserRoutes(app)

	port := os.Getenv("PORT") // Remember to change on Dockerfile
	if strings.TrimSpace(port) == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)

	log.Fatalln(app.Listen(addr))
}
