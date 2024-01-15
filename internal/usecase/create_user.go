package usecase

import (
	"clean-project-exp/internal/entity"
)

type CreateUserUseCase struct {
	WriteUserRepository entity.WriteUserRepositoryInterface
}

func NewCreateUserUseCase(writeUserRepository entity.WriteUserRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{WriteUserRepository: writeUserRepository}
}

func (u *CreateUserUseCase) Execute(input CreateUserInputDTO) (CreateUserOutputDTO, error) {
	user := entity.NewUser(input.Name)

	if err := u.WriteUserRepository.Create(user); err != nil {
		return CreateUserOutputDTO{}, err
	}

	output := CreateUserOutputDTO{
		PubId:     user.PubId,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return output, nil
}
