package factory

import (
	"clean-project-exp/internal/entity"
	"clean-project-exp/internal/infra/database"
	"clean-project-exp/internal/infra/database/memory"
	"clean-project-exp/internal/usecase"
)

func makeCreateUserUseCase(db database.DBInterface[entity.User]) *usecase.CreateUserUseCase {
	return usecase.NewCreateUserUseCase(memory.NewWriteUserRepository(db))
}

func makeListUsersUseCase(db database.DBInterface[entity.User]) *usecase.ListUsersUseCase {
	return usecase.NewListUsersUseCase(memory.NewListUsersRepository(db))
}
