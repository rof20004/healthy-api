package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	appHandler "github.com/rof20004/healthy-api/adapters/ui/gofiber/handlers/app"
	consultaHandler "github.com/rof20004/healthy-api/adapters/ui/gofiber/handlers/consulta"
	pacienteHandler "github.com/rof20004/healthy-api/adapters/ui/gofiber/handlers/paciente"
	profissionalHandler "github.com/rof20004/healthy-api/adapters/ui/gofiber/handlers/profissional"
	profissionalAgendaHandler "github.com/rof20004/healthy-api/adapters/ui/gofiber/handlers/profissional_agenda"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title Ajuda API
// @version 1.0.0
// @description Sistema de agendamento de consultas médicas para a prefeitura de Embu-Guaçu
// @termsOfService http://swagger.io/terms/
// @contact.name Rodolfo do Nascimento Azevedo
// @contact.email rof20004@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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
