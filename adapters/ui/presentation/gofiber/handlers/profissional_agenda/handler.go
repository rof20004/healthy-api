package profissional_agenda

import (
	"github.com/rof20004/healthy-api/adapters/infrastructure/database/postgres"
	profissionalAgendaPostgreSQLPersistenceAdapter "github.com/rof20004/healthy-api/adapters/infrastructure/database/postgres/profissional_agenda"
	"github.com/rof20004/healthy-api/adapters/ui/presentation/gofiber/handlers"
	profissionalAgendaInputPort "github.com/rof20004/healthy-api/application/domains/profissional_agenda/ports/input"
	profissionalAgendaServices "github.com/rof20004/healthy-api/application/domains/profissional_agenda/services"

	"github.com/gofiber/fiber/v2"
)

var (
	createProfissionalAgendaInputPort  profissionalAgendaInputPort.CreateProfissionalAgendaInputPort
	getAllProfissionalAgendasInputPort profissionalAgendaInputPort.GetAllProfissionalAgendasInputPort
)

func init() {
	postgresDb := postgres.GetInstance()

	profissionalAgendaRepo := profissionalAgendaPostgreSQLPersistenceAdapter.NewPostgreSQLProfissionalAgendaPersistenceAdapter(postgresDb)

	createProfissionalAgendaInputPort = profissionalAgendaServices.NewCreateProfissionalAgendaService(profissionalAgendaRepo)
	getAllProfissionalAgendasInputPort = profissionalAgendaServices.NewGetAllProfissionalAgendaService(profissionalAgendaRepo)
}

func InitProfissionalAgendaRoutes(app *fiber.App) {
	app.Post("/profissional-agendas", create)
	app.Get("/profissional-agendas/profissionais/:profissionalId", getAll)
}

func create(c *fiber.Ctx) error {
	var payload CreateProfissionalAgendaRequest
	if err := c.BodyParser(&payload); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusUnprocessableEntity, err)
	}

	profissionalAgenda := payload.ToEntity()

	if err := createProfissionalAgendaInputPort.CreateProfissionalAgenda(profissionalAgenda); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, profissionalAgenda, fiber.StatusCreated, nil)
}

func getAll(c *fiber.Ctx) error {
	profissionalAgenda, err := getAllProfissionalAgendasInputPort.GetAllProfissionalAgendasByProfissionalId(c.Params("profissionalId"))
	if err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, profissionalAgenda, fiber.StatusOK, nil)
}
