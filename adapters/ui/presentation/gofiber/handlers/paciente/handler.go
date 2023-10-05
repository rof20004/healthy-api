package paciente

import (
	"github.com/rof20004/healthy-api/adapters/infrastructure/database/postgres"
	pacientePostgreSQLPersistenceAdapter "github.com/rof20004/healthy-api/adapters/infrastructure/database/postgres/paciente"
	"github.com/rof20004/healthy-api/adapters/ui/presentation/gofiber/handlers"
	pacienteInputPort "github.com/rof20004/healthy-api/application/domains/paciente/ports/input"
	pacienteServices "github.com/rof20004/healthy-api/application/domains/paciente/services"

	"github.com/gofiber/fiber/v2"
)

var (
	createPacienteInputPort  pacienteInputPort.CreatePacienteInputPort
	getAllPacientesInputPort pacienteInputPort.GetAllPacientesInputPort
)

func init() {
	postgresDb := postgres.GetInstance()

	pacienteRepo := pacientePostgreSQLPersistenceAdapter.NewPostgreSQLPacientePersistenceAdapter(postgresDb)

	createPacienteInputPort = pacienteServices.NewCreatePacienteService(pacienteRepo)
	getAllPacientesInputPort = pacienteServices.NewGetAllPacientesService(pacienteRepo)
}

func InitPacienteRoutes(app *fiber.App) {
	app.Post("/pacientes", create)
	app.Get("/pacientes", getAll)
}

// @Tags		Paciente
// @Summary		Cria um novo paciente
// @Accept		json
// @Produce		json
// @Param		data body CreatePacienteRequest true "Informações do paciente"
// @Success		201
// @Failure		400
// @Failure		500
// @Router		/pacientes [post]
func create(c *fiber.Ctx) error {
	var payload CreatePacienteRequest
	if err := c.BodyParser(&payload); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusUnprocessableEntity, err)
	}

	paciente := payload.ToEntity()

	if err := createPacienteInputPort.CreatePaciente(paciente); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, paciente, fiber.StatusCreated, nil)
}

func getAll(c *fiber.Ctx) error {
	pacientes, err := getAllPacientesInputPort.GetAllPacientes()
	if err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, pacientes, fiber.StatusOK, nil)
}
