package usecase

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserInputDTO struct {
	Name string `json:"name"`
}

type CreateUserOutputDTO struct {
	PubId     uuid.UUID `json:"pub_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserForListingDTO struct {
	PubId     uuid.UUID `json:"pub_id" xml:"pub_id"`
	Name      string    `json:"name" xml:"name"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
}

type ListUsersOutputDTO struct {
	Users []UserForListingDTO `json:"users" xml:"users"`
}
