package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	appHandler "github.com/rof20004/healthy-api/adapters/ui/presentation/gofiber/handlers/app"
	consultaHandler "github.com/rof20004/healthy-api/adapters/ui/presentation/gofiber/handlers/consulta"
	pacienteHandler "github.com/rof20004/healthy-api/adapters/ui/presentation/gofiber/handlers/paciente"
	profissionalHandler "github.com/rof20004/healthy-api/adapters/ui/presentation/gofiber/handlers/profissional"
	profissionalAgendaHandler "github.com/rof20004/healthy-api/adapters/ui/presentation/gofiber/handlers/profissional_agenda"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	appHandler.InitAppRoutes(app)
	consultaHandler.InitConsultaRoutes(app)
	pacienteHandler.InitPacienteRoutes(app)
	profissionalHandler.InitProfissionalRoutes(app)
	profissionalAgendaHandler.InitProfissionalAgendaRoutes(app)

	port := os.Getenv("PORT") // Remember to change on Dockerfile
	if strings.TrimSpace(port) == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)

	log.Fatalln(app.Listen(addr))
}
