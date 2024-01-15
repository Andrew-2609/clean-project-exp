package factory

import (
	"clean-project-exp/internal/infra/database/memory"
	"clean-project-exp/internal/infra/web/handler"
)

func MakeWebUserHandler(db *memory.InMemoryDB) *handler.UserWebHandler {
	return handler.NewUserWebHandler(makeCreateUserUseCase(db), makeListUsersUseCase(db))
}
