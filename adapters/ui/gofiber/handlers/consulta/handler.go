package consulta

import (
	"github.com/rof20004/healthy-api/adapters/database/postgres"
	consultaPostgreSQLPersistenceAdapter "github.com/rof20004/healthy-api/adapters/database/postgres/consulta"
	"github.com/rof20004/healthy-api/adapters/ui/gofiber/handlers"
	consultaInputPort "github.com/rof20004/healthy-api/application/domains/consulta/ports/input"
	consultaServices "github.com/rof20004/healthy-api/application/domains/consulta/services"

	"github.com/gofiber/fiber/v2"
)

var (
	createConsultaInputPort  consultaInputPort.CreateConsultaInputPort
	getAllConsultasInputPort consultaInputPort.GetAllConsultasInputPort
)

func init() {
	postgresDb := postgres.GetInstance()

	consultaRepo := consultaPostgreSQLPersistenceAdapter.NewPostgreSQLConsultaPersistenceAdapter(postgresDb)

	createConsultaInputPort = consultaServices.NewCreateConsultaService(consultaRepo)
	getAllConsultasInputPort = consultaServices.NewGetAllProfissionaisService(consultaRepo)
}

func InitConsultaRoutes(app *fiber.App) {
	app.Post("/consultas", create)
	app.Get("/consultas", getAll)
	app.Get("/consultas/:pacienteId", getAllByPaciente)
}

// @Tags		Consulta
// @Summary		Cria uma consulta para um paciente
// @Accept		json
// @Produce		json
// @Param		data body CreateConsultaRequest true "Informações do paciente, do profissional e a data da consulta"
// @Success		200
// @Failure		400
// @Failure		500
// @Router		/consultas [post]
func create(c *fiber.Ctx) error {
	var payload CreateConsultaRequest
	if err := c.BodyParser(&payload); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusUnprocessableEntity, err)
	}

	consulta := payload.ToEntity()

	if err := createConsultaInputPort.CreateConsulta(consulta); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, consulta, fiber.StatusCreated, nil)
}

// @Tags		Consulta
// @Summary		Retorna a lista de consultas de todos os pacientes
// @Produce		json
// @Success		200
// @Failure		400
// @Failure		500
// @Router		/consultas [get]
func getAll(c *fiber.Ctx) error {
	consultas, err := getAllConsultasInputPort.GetAllConsultas()
	if err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, consultas, fiber.StatusOK, nil)
}

// @Tags		Consulta
// @Summary		Retorna a lista de consultas de um determinado paciente
// @Produce		json
// @Param		pacienteId path string true "Id do paciente"
// @Success		200
// @Failure		400
// @Failure		500
// @Router		/consultas/{pacienteId} [get]
func getAllByPaciente(c *fiber.Ctx) error {
	consultas, err := getAllConsultasInputPort.GetAllConsultasByPacienteId(c.Params("pacienteId"))
	if err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, consultas, fiber.StatusOK, nil)
}
