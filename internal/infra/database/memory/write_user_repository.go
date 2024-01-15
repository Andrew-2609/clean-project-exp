package memory

import (
	"clean-project-exp/internal/entity"
	"clean-project-exp/internal/infra/database"
	"github.com/google/uuid"
	"time"
)

type WriteUserRepository struct {
	db database.DBInterface[entity.User]
}

func NewWriteUserRepository(db database.DBInterface[entity.User]) *WriteUserRepository {
	return &WriteUserRepository{db: db}
}

func (r *WriteUserRepository) Create(user *entity.User) error {
	user.ID = r.db.Count() + 1
	user.PubId = uuid.New()
	user.CreatedAt = time.Now()
	return r.db.Save(user)
}
