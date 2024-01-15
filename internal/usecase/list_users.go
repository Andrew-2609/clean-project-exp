package usecase

import (
	"clean-project-exp/internal/entity"
)

type ListUsersUseCase struct {
	ListUsersRepository entity.ListUsersRepositoryInterface
}

func NewListUsersUseCase(listUsersRepository entity.ListUsersRepositoryInterface) *ListUsersUseCase {
	return &ListUsersUseCase{ListUsersRepository: listUsersRepository}
}

func (u *ListUsersUseCase) Execute() (ListUsersOutputDTO, error) {
	usersFromDB := u.ListUsersRepository.List()
	var usersForListing []UserForListingDTO

	for _, user := range usersFromDB {
		usersForListing = append(usersForListing, UserForListingDTO{
			PubId:     user.PubId,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
		})
	}

	output := ListUsersOutputDTO{Users: usersForListing}

	return output, nil
}
