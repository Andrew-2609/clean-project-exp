package memory

import (
	"clean-project-exp/internal/entity"
)

type InMemoryDB struct {
	users []*entity.User
}

func InitInMemoryDB() *InMemoryDB {
	return &InMemoryDB{users: make([]*entity.User, 0)}
}

func (db *InMemoryDB) List() []*entity.User {
	return db.users
}

func (db *InMemoryDB) Count() int64 {
	return int64(len(db.users))
}

func (db *InMemoryDB) Save(user *entity.User) error {
	db.users = append(db.users, user)
	return nil
}
