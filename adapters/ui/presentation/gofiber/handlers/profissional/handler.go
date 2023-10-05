package profissional

import (
	"github.com/rof20004/healthy-api/adapters/infrastructure/database/postgres"
	profissionalPostgreSQLPersistenceAdapter "github.com/rof20004/healthy-api/adapters/infrastructure/database/postgres/profissional"
	"github.com/rof20004/healthy-api/adapters/ui/presentation/gofiber/handlers"
	profissionalInputPort "github.com/rof20004/healthy-api/application/domains/profissional/ports/input"
	profissionalServices "github.com/rof20004/healthy-api/application/domains/profissional/services"

	"github.com/gofiber/fiber/v2"
)

var (
	createProfissionalInputPort  profissionalInputPort.CreateProfissionalInputPort
	getAllProfissionaisInputPort profissionalInputPort.GetAllProfissionaisInputPort
)

func init() {
	postgresDb := postgres.GetInstance()

	profissionalRepo := profissionalPostgreSQLPersistenceAdapter.NewPostgreSQLProfissionalPersistenceAdapter(postgresDb)

	createProfissionalInputPort = profissionalServices.NewCreateProfissionalService(profissionalRepo)
	getAllProfissionaisInputPort = profissionalServices.NewGetAllProfissionaisService(profissionalRepo)
}

func InitProfissionalRoutes(app *fiber.App) {
	app.Post("/profissionais", create)
	app.Get("/profissionais", getAll)
}

func create(c *fiber.Ctx) error {
	var payload CreateProfissionalRequest
	if err := c.BodyParser(&payload); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusUnprocessableEntity, err)
	}

	profissional := payload.ToEntity()

	if err := createProfissionalInputPort.CreateProfissionalAgenda(profissional); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, profissional, fiber.StatusCreated, nil)
}

func getAll(c *fiber.Ctx) error {
	profissionais, err := getAllProfissionaisInputPort.GetAllProfissionais()
	if err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, profissionais, fiber.StatusOK, nil)
}
