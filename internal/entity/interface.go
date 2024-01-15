package entity

type WriteUserRepositoryInterface interface {
	Create(user *User) error
}

type ListUsersRepositoryInterface interface {
	List() []User
}
