package memory

import (
	"clean-project-exp/internal/entity"
	"clean-project-exp/internal/infra/database"
)

type ListUsersRepository struct {
	db database.DBInterface[entity.User]
}

func NewListUsersRepository(db database.DBInterface[entity.User]) *ListUsersRepository {
	return &ListUsersRepository{db: db}
}

func (r *ListUsersRepository) List() []entity.User {
	var output []entity.User

	for _, user := range r.db.List() {
		if user != nil {
			output = append(output, *user)
		}
	}

	return output
}
