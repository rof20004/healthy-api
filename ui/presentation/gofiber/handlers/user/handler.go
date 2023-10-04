package user

import (
	userInputPort "github.com/rof20004/healthy-api/application/domains/user/ports/input"
	userServices "github.com/rof20004/healthy-api/application/domains/user/services"
	"github.com/rof20004/healthy-api/infrastructure/database/postgres"
	userPostgreSQLPersistenceAdapter "github.com/rof20004/healthy-api/infrastructure/database/postgres/user"
	"github.com/rof20004/healthy-api/ui/presentation/gofiber/handlers"

	"github.com/gofiber/fiber/v2"
)

var (
	createUserInputPort  userInputPort.CreateUserInputPort
	getAllUsersInputPort userInputPort.GetAllUsersInputPort
)

func init() {
	postgresDb := postgres.GetInstance()

	userRepo := userPostgreSQLPersistenceAdapter.NewPostgreSQLUserPersistenceAdapter(postgresDb)

	createUserInputPort = userServices.NewCreateUserService(userRepo)
	getAllUsersInputPort = userServices.NewGetAllUsersService(userRepo)
}

func InitUserRoutes(app *fiber.App) {
	app.Post("/users", create)
	app.Get("/users", getAll)
}

func create(c *fiber.Ctx) error {
	var payload CreateUserRequest
	if err := c.BodyParser(&payload); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusUnprocessableEntity, err)
	}

	user := payload.ToEntity()

	if err := createUserInputPort.CreateUser(user); err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, user, fiber.StatusCreated, nil)
}

func getAll(c *fiber.Ctx) error {
	users, err := getAllUsersInputPort.GetAllUsers()
	if err != nil {
		return handlers.SendResponse(c, nil, fiber.StatusBadRequest, err)
	}

	return handlers.SendResponse(c, users, fiber.StatusOK, nil)
}
