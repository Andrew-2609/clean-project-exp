package database

import "clean-project-exp/internal/entity"

type DBInterface[T entity.User] interface {
	Save(entity *T) error
	List() []*T
	Count() int64
}
